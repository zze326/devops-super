package user

import (
	"context"
	"devops-super/internal/consts"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"
	"devops-super/utility/thirdclients/kubernetes"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

func (s *sCiPipeline) Run(ctx context.Context, id int) (err error) {
	var (
		ePipeline        *entity.CiPipeline
		config           mid.CiPipelineConfig
		kubeConfigSecret *entity.Secret
		kubeConfig       *mid.TextContent
		kubeClient       *kubernetes.Client
		envMap           map[int]*entity.CiEnv
	)
	ePipeline, err = s.Get(ctx, &do.CiPipeline{Id: id})
	if err != nil {
		return
	}

	if err = ePipeline.Config.Scan(&config); err != nil {
		return
	}

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

	if kubeConfigSecret, err = service.Secret().Get(ctx, &do.Secret{Id: ePipeline.KubernetesConfigId}); err != nil {
		return
	}

	if err = kubeConfigSecret.Content.Scan(&kubeConfig); err != nil {
		return
	}
	if kubeClient, err = kubernetes.NewClient(kubeConfig.Text); err != nil {
		return
	}

	return createCiPod(kubeClient, fmt.Sprintf("ci-pod-%d-%s", ePipeline.Id, time.Now().Format("20060102150405")), config)
}

// 创建 ci pod
func createCiPod(kubeClient *kubernetes.Client, name string, ciConfig mid.CiPipelineConfig) error {
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
			Namespace: consts.CI_CLIENT_POD_NAMESPACE,
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
	if _, err := kubeClient.CoreV1().Pods(consts.CI_CLIENT_POD_NAMESPACE).Create(context.Background(), pod, metav1.CreateOptions{}); err != nil {
		return err
	}
	return nil
}
