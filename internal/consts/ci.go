package consts

// 流水线任务类型
const (
	PIPELINE_TASK_TYPE_GIT_PULL   = 1 // Git 拉取
	PIPELINE_TASK_TYPE_SHELL_EXEC = 2 // 执行 shell
)

// CI 客户端 Pod 相关
const (
	CI_CLIENT_POD_WORKSPACE_VOLUME_NAME     = "ci-workspace"    // ci 客户端 pod 工作卷名称
	CI_CLIENT_POD_WORKSPACE_PATH            = "/devops-super"   // ci 客户端 pod 工作卷挂载目录
	CI_CLIENT_POD_NAMESPACE                 = "devops-super-ci" // ci 客户端 pod 所在命名空间
	CI_CLIENT_POD_CONTAINER_STAGES_ENV_NAME = "STAGES"          // ci 客户端容器阶段环境信息环境变量名称
)
