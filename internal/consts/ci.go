package consts

// 流水线任务类型
const (
	PIPELINE_TASK_TYPE_GIT_PULL   = 1 // Git 拉取
	PIPELINE_TASK_TYPE_SHELL_EXEC = 2 // 执行 shell
)

// CI 客户端 Pod 相关
const (
	CI_CLIENT_POD_WORKSPACE_VOLUME_NAME             = "ci-workspace"                                                       // ci 客户端 pod 工作卷名称
	CI_CLIENT_POD_WORKSPACE_PATH                    = "/devops-super"                                                      // ci 客户端 pod 工作卷挂载目录
	CI_CLIENT_POD_NAMESPACE                         = "devops-super-ci"                                                    // ci 客户端 pod 所在命名空间
	CI_CLIENT_POD_CONTAINER_STAGES_ENV_NAME         = "STAGES"                                                             // ci 客户端容器阶段环境信息环境变量名称
	CI_CLIENT_POD_KANIKO_DIR                        = "/kaniko"                                                            // ci 客户端 kaniko 环境目录
	CI_CLIENT_POD_KANIKO_WORKSPACE_PATH             = "/workspace"                                                         // ci 客户端 kaniko 容器的工作目录
	CI_CLIENT_POD_KANIKO_EXECUTOR_IMAGE             = "registry.cn-shenzhen.aliyuncs.com/zze/gcriokaniko-executor:v1.19.0" // kaniko executor 镜像
	CI_CLIENT_POD_KANIKO_WARMER_IMAGE               = "registry.cn-shenzhen.aliyuncs.com/zze/gcriokaniko-warmer:v1.19.0"   // kanico 推送地址
	CI_CLIENT_POD_KANIKO_IMAGE_DESTINATION_ENV_NAME = "IMAGE_DESTINATION"                                                  // kanico warmer 镜像
	CI_CLIENT_CONFIG_MAP_NAME                       = "devops-super-ci-config"
	CI_CLIENT_CONFIG_MAP_SECRET_KEY_PREFIX          = "secret"
)
