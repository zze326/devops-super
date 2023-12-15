package user

import (
	"context"
	"database/sql"
	"devops-super/internal/consts"
	"devops-super/internal/dao"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"
	"devops-super/utility/thirdclients/kubernetes"
	"devops-super/utility/util"
	"encoding/base64"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"path/filepath"
	"time"
)

func (s *sCiPipeline) Run(ctx context.Context, id int, params *gjson.Json) (err error) {
	var (
		ePipeline         *entity.CiPipeline
		arrangeConfigJson *gjson.Json
		arrangeConfig     mid.CiPipelineConfig // 编排配置
		kubeConfigSecret  *entity.Secret
		kubeConfig        *mid.TextContent
		kubeClient        *kubernetes.Client
		kubeNamespace     = consts.CI_CLIENT_POD_NAMESPACE
		envMap            map[int]*entity.CiEnv
		now               = gtime.Now()
	)
	// 获取 Pipeline 信息
	ePipeline, err = s.Get(ctx, &do.CiPipeline{Id: id})
	if err != nil {
		return
	}
	arrangeConfigJson = ePipeline.Config

	if ePipeline.Parameterize {
		if params.IsNil() {
			return gerror.New("参数化流水线未提交参数")
		}

		now := gtime.Now()

		env := g.Map{
			"params": params.Map(),
			"now": g.Map{
				"year":        now.Year(),
				"month":       now.Month(),
				"day":         now.Day(),
				"hour":        now.Hour(),
				"minute":      now.Minute(),
				"second":      now.Second(),
				"millisecond": now.Millisecond(),
				"string1":     now.String(),
				"string2":     now.Layout("20060102150405"),
			},
		}

		arrangeConfigJsonStr, err := util.Pongo2Parse(ePipeline.Config.String(), env)
		if err != nil {
			return err
		}
		arrangeConfigJson = gjson.New(arrangeConfigJsonStr)
	}

	// 解析 Pipeline 编排配置到结构体对象
	if err = arrangeConfigJson.Scan(&arrangeConfig); err != nil {
		return
	}

	// 获取环境信息 map
	if envMap, err = service.CiEnv().GetEntityMap(ctx, arrangeConfig.GetEnvIds()); err != nil {
		return
	}

	// 组装环境关联的镜像和秘钥名称
	for _, envItem := range arrangeConfig {
		envItem.Image = envMap[envItem.Id].Image
		envItem.SecretName = envMap[envItem.Id].SecretName
		envItem.IsKaniko = envMap[envItem.Id].IsKaniko

		for _, stageItem := range envItem.Stages {
			for _, taskItem := range stageItem.Tasks {
				// 如果设置了 git 认证秘钥，组装一下
				if secretId := taskItem.GitPullData.SecretId; secretId > 0 {
					var (
						eSecret      *entity.Secret
						gitBasicAuth *mid.UsernamePasswordContent
					)
					eSecret, err = service.Secret().Get(ctx, &do.Secret{Id: secretId})
					if err != nil {
						return err
					}

					if err = eSecret.Content.Scan(&gitBasicAuth); err != nil {
						return
					}
					if eSecret.Type == consts.SECRET_TYPE_GIT_BASIC_AUTH {
						taskItem.GitPullData.GitBasicAuth = gitBasicAuth
					}
				}
			}
		}
	}

	// 获取 k8s 配置信息
	if kubeConfigSecret, err = service.Secret().Get(ctx, &do.Secret{Id: ePipeline.KubernetesConfigId}); err != nil {
		return
	}
	// 解析 k8s 配置信息
	if err = kubeConfigSecret.Content.Scan(&kubeConfig); err != nil {
		return
	}
	// 创建 k8s 客户端
	if kubeClient, err = kubernetes.NewClient(ctx, kubeConfig.Text); err != nil {
		return gerror.Wrap(err, "Kubernetes Config 无效，客户端连接失败")
	}

	if !gutil.IsEmpty(ePipeline.KubernetesNamespace) {
		kubeNamespace = ePipeline.KubernetesNamespace
	}

	ciPodName := fmt.Sprintf("ci-%s-%d-%s", ePipeline.Name, id, time.Now().Format("20060102150405"))

	// 创建 ci pod
	if err = createCiPod(kubeClient, kubeNamespace, ciPodName, arrangeConfig, envMap); err != nil {
		return
	}

	var (
		r     sql.Result
		runId int64
	)
	// 数据库插入 pipeline 运行记录
	r, err = dao.CiPipelineRun.Ctx(ctx).Insert(&entity.CiPipelineRun{
		PipelineId: ePipeline.Id,
		PodName:    ciPodName,
		Namespace:  kubeNamespace,
		Status:     0,
		UpdatedAt:  now,
		CreatedAt:  now,
	})

	if err != nil {
		return
	}

	runId, err = r.LastInsertId()
	if err != nil {
		return
	}

	// 协程监听 pod 状态
	go watchCiPod(g.RequestFromCtx(ctx).GetNeverDoneCtx(), kubeClient, kubeNamespace, ciPodName, int(runId))
	return
}

// 监听 ci pod
func watchCiPod(ctx context.Context, kubeClient *kubernetes.Client, namespace, name string, runId int) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error(ctx, err)
		}
	}()
	watcher, err := kubeClient.CoreV1().Pods(namespace).Watch(ctx, metav1.ListOptions{
		FieldSelector: fmt.Sprintf("metadata.name=%s", name),
	})
	if err != nil {
		glog.Errorf(ctx, "Failed to create watcher, err: %v", err)
		return
	}
	defer watcher.Stop()

WATCH:
	for event := range watcher.ResultChan() {
		switch event.Type {
		case watch.Modified:
			var pod = event.Object.(*corev1.Pod)
			if pod.Status.Phase == corev1.PodPending || pod.Status.Phase == corev1.PodRunning {
			} else if pod.Status.Phase == corev1.PodSucceeded {
				glog.Debugf(ctx, "Pod '%s' modified in namespace '%s' Success", pod.Name, pod.Namespace)
				if _, err := dao.CiPipelineRun.Ctx(ctx).WherePri(runId).Data(do.CiPipelineRun{
					Status: 1,
				}).OmitNilData().Update(); err != nil {
					glog.Error(ctx, err)
					return
				}

				break WATCH
			} else if pod.Status.Phase == corev1.PodFailed {
				glog.Debugf(ctx, "Pod '%s' modified in namespace '%s' Failed", pod.Name, pod.Namespace)
				cols := dao.CiPipelineRun.Columns()
				if _, err := dao.CiPipelineRun.Ctx(ctx).Where(g.Map{
					cols.Id:     runId,
					cols.Status: 0,
				}).Data(do.CiPipelineRun{
					Status: 2,
				}).OmitNilData().Update(); err != nil {
					glog.Error(ctx, err)
					return
				}
				break WATCH
			}
		case watch.Error:
			//err := event.Object.(error)
			glog.Errorf(ctx, "Received watch error: %v", event.Object)
			break WATCH
		}
	}
}

// 创建 ci pod
func createCiPod(kubeClient *kubernetes.Client, namespace, name string, arrangeConfig mid.CiPipelineConfig, envMap map[int]*entity.CiEnv) error {
	var (
		containers              []corev1.Container
		initContainers          []corev1.Container
		volumes                 []corev1.Volume
		imagePullSecretNames    = gset.New()
		dockerRegistrySecretIds = gset.New()
		imagePullSecrets        []corev1.LocalObjectReference
		hasKaniko               bool
		dockerfilePathsToCache  []string
		kanikoVolumeMounts      []corev1.VolumeMount
	)

	var createEnvs = func(envs map[string]string) []corev1.EnvVar {
		var result []corev1.EnvVar
		for k, v := range envs {
			result = append(result, corev1.EnvVar{Name: k, Value: v})
		}
		return result
	}

	for idx, envItem := range arrangeConfig {
		mountPath := consts.CI_CLIENT_POD_WORKSPACE_PATH
		containerName := fmt.Sprintf("env-%d", idx)
		if envItem.IsKaniko {
			hasKaniko = true
			mountPath = consts.CI_CLIENT_POD_KANIKO_WORKSPACE_PATH
			containerName = fmt.Sprintf("%s-kaniko", containerName)
			dockerRegistrySecretIds.Add(envItem.KanikoParam.SecretId)
		}
		container := corev1.Container{
			Name:            containerName,
			Image:           envItem.Image,
			ImagePullPolicy: corev1.PullAlways,
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      consts.CI_CLIENT_POD_WORKSPACE_VOLUME_NAME,
					MountPath: mountPath,
				},
			},
		}
		if envItem.IsKaniko {
			container.Image = consts.CI_CLIENT_POD_KANIKO_EXECUTOR_IMAGE
			container.Args = append(container.Args, "--cache=true")
			container.Args = append(container.Args, fmt.Sprintf("--kaniko-dir=%s", consts.CI_CLIENT_POD_KANIKO_DIR))
			container.Args = append(container.Args, "--cache-dir=/cache")
			container.Args = append(container.Args, "--skip-tls-verify")
			container.Args = append(container.Args, "--skip-tls-verify-pull")
			container.Args = append(container.Args, fmt.Sprintf("--context=dir://%s", envItem.KanikoParam.ContextDir))
			container.Args = append(container.Args, fmt.Sprintf("--destination=%s", envItem.KanikoParam.ImageDestination))
			if envItem.KanikoParam.UpdateBaseImageCache {
				dockerfilePathsToCache = append(dockerfilePathsToCache, envItem.KanikoParam.DockerfilePath)
			}

			container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{
				Name:      consts.CI_CLIENT_CONFIG_MAP_NAME,
				ReadOnly:  true,
				MountPath: filepath.Join(consts.CI_CLIENT_POD_KANIKO_DIR, "/.docker/config.json"),
				SubPath:   fmt.Sprintf("%s-%d", consts.CI_CLIENT_CONFIG_MAP_SECRET_KEY_PREFIX, envItem.KanikoParam.SecretId),
			})
		} else {
			stagesJson, err := gjson.EncodeString(envItem.Stages)
			if err != nil {
				return err
			}
			container.Env = createEnvs(map[string]string{
				consts.CI_CLIENT_POD_CONTAINER_STAGES_ENV_NAME: stagesJson,
			})
		}

		// 判断该环境启用了持久化
		var persistenceConfig mid.CiEnvPersistenceConfig
		if persistenceConfigJson := envMap[envItem.Id].PersistenceConfig; !persistenceConfigJson.IsNil() {
			if err := envMap[envItem.Id].PersistenceConfig.Scan(&persistenceConfig); err != nil {
				return err
			}
		}
		for _, item := range persistenceConfig {
			volumeName := item.PvcName
			existsPvc := false
			for _, volume := range volumes {
				if item.PvcName == volume.Name {
					existsPvc = true
					break
				}
			}
			if !existsPvc {
				volumes = append(volumes, corev1.Volume{
					Name: volumeName,
					VolumeSource: corev1.VolumeSource{
						PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
							ClaimName: item.PvcName,
						},
					},
				})
			}

			container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{
				Name:      volumeName,
				MountPath: item.MountPath,
				SubPath:   item.SubPath,
			})
		}
		if envItem.IsKaniko {
			kanikoVolumeMounts = container.VolumeMounts
		}

		updateBaseImageCache := hasKaniko && len(dockerfilePathsToCache) > 0

		if len(arrangeConfig) == 1 || idx == (len(arrangeConfig)-1) { // 如果只有一个环境容器 或 是最后一个容器
			if updateBaseImageCache { // 如果要更新基础镜像缓存，则最后一个容器肯定是 kaniko warm 容器
				initContainers = append(initContainers, container)
			} else {
				containers = append(containers, container)
			}
		} else {
			initContainers = append(initContainers, container)
		}

		// 如果存在 kaniko 环境，缓存构建使用的镜像
		if updateBaseImageCache {
			cacheContainer := corev1.Container{
				Name:            fmt.Sprintf("env-%d-kaniko-warmer", len(arrangeConfig)),
				Image:           consts.CI_CLIENT_POD_KANIKO_WARMER_IMAGE,
				ImagePullPolicy: corev1.PullIfNotPresent,
				VolumeMounts:    kanikoVolumeMounts,
			}
			cacheContainer.Args = append(cacheContainer.Args, "--cache-dir=/cache")
			cacheContainer.Args = append(cacheContainer.Args, "--skip-tls-verify-pull")
			cacheContainer.Args = append(cacheContainer.Args, "--force")
			//cacheContainer.Args = append(cacheContainer.Args, fmt.Sprintf("--image=%s", "registry-azj-registry.cn-shanghai.cr.aliyuncs.com/ops/alpine-for-dp:v1.2"))
			for _, dockerfilePath := range dockerfilePathsToCache {
				cacheContainer.Args = append(cacheContainer.Args, fmt.Sprintf("--dockerfile=%s", dockerfilePath))
			}
			containers = append(containers, cacheContainer)
		}

		if !gutil.IsEmpty(envItem.SecretName) {
			imagePullSecretNames.Add(envItem.SecretName)
		}
	}

	for _, imageSecretName := range imagePullSecretNames.Slice() {
		imagePullSecrets = append(imagePullSecrets, corev1.LocalObjectReference{Name: imageSecretName.(string)})
	}

	if hasKaniko { // 如果存在 kaniko 环境，检查 kaniko 环境配置的镜像仓库认证秘钥内容和 kubernetes 集群中 ConfigMap 中的配置是否一致
		configMap, err := kubeClient.GetConfigMap(namespace, consts.CI_CLIENT_CONFIG_MAP_NAME)
		if err != nil && !kubernetes.IsNotFoundError(err) {
			return err
		}

		var (
			noConfigMap           = kubernetes.IsNotFoundError(err)
			configDataMap         = make(map[string]string)
			shouldUpdateConfigMap = false
		)

		for _, secretId := range dockerRegistrySecretIds.Slice() {
			eSecret, err := service.Secret().Get(kubeClient.Ctx, &do.Secret{Id: secretId.(int)})
			if err != nil {
				return err
			}

			if eSecret.Type != consts.SECRET_TYPE_DOCKER_REGISTRY_AUTH {
				return gerror.New("秘钥类型不匹配")
			}
			dockerRegistryAuthContent := new(mid.DockerRegistryAuthContent)
			if err := eSecret.Content.Scan(dockerRegistryAuthContent); err != nil {
				return err
			}

			authConfigJson, err := gjson.MarshalIndent(g.Map{
				"auths": g.Map{
					dockerRegistryAuthContent.RegistryUrl: g.Map{
						"auth": base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", dockerRegistryAuthContent.Username, dockerRegistryAuthContent.Password))),
					},
				},
			}, "", "  ")
			if err != nil {
				return err
			}

			itemKey := fmt.Sprintf("%s-%d", consts.CI_CLIENT_CONFIG_MAP_SECRET_KEY_PREFIX, secretId)
			configDataMap[fmt.Sprintf(itemKey)] = string(authConfigJson)
			if !noConfigMap { // 存在 config map
				if content, ok := configMap.Data[itemKey]; !ok || content != string(authConfigJson) {
					shouldUpdateConfigMap = true
					configMap.Data[itemKey] = string(authConfigJson)
				}
			}
		}

		if noConfigMap {
			if err := kubeClient.CreateConfigMap(namespace, consts.CI_CLIENT_CONFIG_MAP_NAME, configDataMap); err != nil {
				return err
			}
		} else {
			if shouldUpdateConfigMap {
				if err := kubeClient.UpdateConfigMap(namespace, configMap); err != nil {
					return err
				}
			}
		}

		volumes = append(volumes, corev1.Volume{
			Name: consts.CI_CLIENT_CONFIG_MAP_NAME,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: consts.CI_CLIENT_CONFIG_MAP_NAME,
					},
				},
			},
		})
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: corev1.PodSpec{
			ImagePullSecrets: imagePullSecrets,
			InitContainers:   initContainers,
			Containers:       containers,
			Volumes: append(volumes, corev1.Volume{
				Name: consts.CI_CLIENT_POD_WORKSPACE_VOLUME_NAME,
				VolumeSource: corev1.VolumeSource{
					EmptyDir: &corev1.EmptyDirVolumeSource{},
				},
			},
			),
			RestartPolicy: corev1.RestartPolicyNever,
		},
	}

	pvcs, err := kubeClient.GetPersistentVolumeClaims(namespace)
	if err != nil {
		return gerror.Wrap(err, "获取集群 PVC 信息失败")
	}

	for _, volume := range volumes {
		if volume.VolumeSource.PersistentVolumeClaim != nil && !util.InSlice(pvcs, volume.Name) {
			return gerror.Newf("集群的 %s 命名空间下不存在名为 %s 的 PVC", namespace, volume.Name)
		}
	}

	if _, err := kubeClient.CoreV1().Pods(namespace).Create(kubeClient.Ctx, pod, metav1.CreateOptions{}); err != nil {
		return err
	}
	return nil
}
