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
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"time"
)

func (s *sCiPipeline) Run(ctx context.Context, id int) (err error) {
	var (
		ePipeline        *entity.CiPipeline
		config           mid.CiPipelineConfig
		kubeConfigSecret *entity.Secret
		kubeConfig       *mid.TextContent
		kubeClient       *kubernetes.Client
		kubeNamespace    = consts.CI_CLIENT_POD_NAMESPACE
		envMap           map[int]*entity.CiEnv
		now              = gtime.Now()
	)
	// 获取 Pipeline 信息
	ePipeline, err = s.Get(ctx, &do.CiPipeline{Id: id})
	if err != nil {
		return
	}

	// 解析 Pipeline 配置到结构体对象
	if err = ePipeline.Config.Scan(&config); err != nil {
		return
	}

	// 获取环境信息 map
	if envMap, err = service.CiEnv().GetEntityMap(ctx, config.GetEnvIds()); err != nil {
		return
	}

	// 组装环境关联的镜像和秘钥名称
	for _, envItem := range config {
		envItem.Image = envMap[envItem.Id].Image
		envItem.SecretName = envMap[envItem.Id].SecretName

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
	if kubeClient, err = kubernetes.NewClient(kubeConfig.Text); err != nil {
		return
	}

	if !gutil.IsEmpty(ePipeline.KubernetesNamespace) {
		kubeNamespace = ePipeline.KubernetesNamespace
	}

	ciPodName := fmt.Sprintf("ci-%s-%d-%s", ePipeline.Name, id, time.Now().Format("20060102150405"))

	// 创建 ci pod
	if err = createCiPod(kubeClient, kubeNamespace, ciPodName, config); err != nil {
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
	go watchCiPod(g.RequestFromCtx(ctx).GetNeverDoneCtx(), kubeClient, kubeNamespace, ciPodName, config, int(runId))
	return
}

// 监听 ci pod
func watchCiPod(ctx context.Context, kubeClient *kubernetes.Client, namespace, name string, ciConfig mid.CiPipelineConfig, runId int) {
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

	//var (
	//	watchIndex = 0
	//	logIndex   = 0
	//	maxIndex   = len(ciConfig) - 1
	//)
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
				if _, err := dao.CiPipelineRun.Ctx(ctx).WherePri(runId).Data(do.CiPipelineRun{
					Status: 2,
				}).OmitNilData().Update(); err != nil {
					glog.Error(ctx, err)
					return
				}
				break WATCH
			}
			//glog.Printf(ctx, "Pod '%s' modified in namespace '%s'", pod.Name, pod.Namespace)
			//if watchIndex != maxIndex {
			//	for _, status := range pod.Status.InitContainerStatuses {
			//		if containerName := fmt.Sprintf("env-%d", watchIndex); status.Name == containerName {
			//			if status.Ready {
			//				glog.Debugf(ctx, "%s 已完成\n", containerName)
			//				watchIndex++
			//			} else {
			//				if status.State.Running != nil {
			//					glog.Debugf(ctx, "%s 开始运行了\n", containerName)
			//				}
			//				if status.State.Terminated != nil && status.State.Terminated.ExitCode != 0 {
			//					glog.Debugf(ctx, "%s 执行失败\n", containerName)
			//				}
			//			}
			//		}
			//
			//		if containerName := fmt.Sprintf("env-%d", logIndex); status.Name == containerName {
			//			if status.Ready {
			//				tailLog(ctx, kubeClient, namespace, name, logIndex, false) // 打印所有日志
			//				logIndex++
			//			} else {
			//				if status.State.Running != nil {
			//					tailLog(ctx, kubeClient, namespace, name, logIndex, true) // 跟踪打印日志
			//					logIndex++
			//				}
			//			}
			//		}
			//	}
			//} else {
			//	for _, status := range pod.Status.ContainerStatuses {
			//		if containerName := fmt.Sprintf("env-%d", watchIndex); status.Name == containerName {
			//			if terminalState := status.State.Terminated; terminalState != nil {
			//				if terminalState.ExitCode == 0 && terminalState.Reason == "Completed" {
			//					glog.Debugf(ctx, "%s 已完成\n", containerName)
			//				} else if terminalState.ExitCode > 0 {
			//					glog.Debugf(ctx, "%s 执行失败\n", containerName)
			//				}
			//			} else {
			//				if status.State.Running != nil {
			//					glog.Debugf(ctx, "%s 开始运行了\n", containerName)
			//				}
			//			}
			//		}
			//
			//		if containerName := fmt.Sprintf("env-%d", logIndex); status.Name == containerName {
			//			if terminalState := status.State.Terminated; terminalState != nil {
			//				tailLog(ctx, kubeClient, namespace, name, logIndex, false) // 打印所有日志
			//				logIndex++
			//			} else {
			//				if status.State.Running != nil {
			//					tailLog(ctx, kubeClient, namespace, name, logIndex, true) // 跟踪打印日志
			//					logIndex++
			//				}
			//			}
			//		}
			//	}
			//}
		case watch.Error:
			//err := event.Object.(error)
			glog.Errorf(ctx, "Received watch error: %v", event.Object)
			break WATCH
		}
	}
}

// 获取 Pod 日志
//func tailLog(ctx context.Context, kubeClient *kubernetes.Client, namespace, podName string, logIndex int, follow bool) {
//	line := int64(100000)
//req := kubeClient.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{
//	Container: fmt.Sprintf("env-%d", logIndex),
//	Follow:    follow,
//	TailLines: &line,
//})
//stream, err := req.Stream(ctx)
//if err != nil {
//	glog.Error(ctx, err)
//	return
//}
//
//if _, err = io.Copy(os.Stdout, stream); err != nil {
//	glog.Error(ctx, err)
//	return
//}

//defer stream.Close()
//scanner := bufio.NewScanner(stream)
//for scanner.Scan() {
//	fmt.Println(scanner.Text())
//}
//}

// 创建 ci pod
func createCiPod(kubeClient *kubernetes.Client, namespace, name string, ciConfig mid.CiPipelineConfig) error {
	var (
		containers     []corev1.Container
		initContainers []corev1.Container
	)

	var createEnvs = func(envs map[string]string) []corev1.EnvVar {
		var result []corev1.EnvVar
		for k, v := range envs {
			result = append(result, corev1.EnvVar{Name: k, Value: v})
		}
		return result
	}

	for idx, envItem := range ciConfig {
		stagesJson, err := gjson.EncodeString(envItem.Stages)
		if err != nil {
			return err
		}
		container := corev1.Container{
			Name:  fmt.Sprintf("env-%d", idx),
			Image: ciConfig[0].Image,
			Env: createEnvs(map[string]string{
				consts.CI_CLIENT_POD_CONTAINER_STAGES_ENV_NAME: stagesJson,
			}),
			ImagePullPolicy: corev1.PullAlways,
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      consts.CI_CLIENT_POD_WORKSPACE_VOLUME_NAME,
					MountPath: consts.CI_CLIENT_POD_WORKSPACE_PATH,
				},
			},
		}
		if len(ciConfig) == 1 { // 如果只有一个环境容器，则设置该容器到 containers
			containers = append(containers, container)
		} else { // 如果有多个环境容器，则最后一个容器设置到 containers，其它容器设置到 initContainers
			if idx != (len(ciConfig) - 1) { // 如果不是最后一个容器
				initContainers = append(initContainers, container)
			} else { // 是最后一个容器
				containers = append(containers, container)
			}
		}
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: corev1.PodSpec{
			InitContainers: initContainers,
			Containers:     containers,
			Volumes: []corev1.Volume{
				{
					Name: consts.CI_CLIENT_POD_WORKSPACE_VOLUME_NAME,
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				},
			},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	}
	if _, err := kubeClient.CoreV1().Pods(namespace).Create(context.Background(), pod, metav1.CreateOptions{}); err != nil {
		return err
	}
	return nil
}