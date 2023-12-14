package consts

// 权限类型
const (
	PERMISSION_TYPE_DIR  = 1 // 目录
	PERMISSION_TYPE_MENU = 2 // 菜单
	PERMISSION_TYPE_ABLE = 3 // 功能
)

// 系统必须权限名称
const PERMISSION_SYSTEM_REQUIRED_NAME = "system-required"

// 主机终端会话文件默认保存目录，可通过配置 host.terminal.sessionFileDir 修改
const HOST_TERMINAL_SESSION_SAVE_DIRECTORY = "host-sessions"

// 秘钥类型
const (
	SECRET_TYPE_GIT_BASIC_AUTH       = 1 // GIT Basic 认证
	SECRET_TYPE_KUBERNETES_CONFIG    = 2 // Kubernetes config
	SECRET_TYPE_DOCKER_REGISTRY_AUTH = 3 // Docker 镜像仓库认证
)
