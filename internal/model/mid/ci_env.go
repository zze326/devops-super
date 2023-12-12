package mid

import "github.com/gogf/gf/v2/encoding/gjson"

type CiEnv struct {
	Name              string      `v:"required|max-length:30" json:"name"`
	Image             string      `v:"required" json:"image"`
	SecretName        string      `json:"secretName"`
	IsKaniko          bool        `json:"isKaniko"`
	PersistenceConfig *gjson.Json `json:"persistenceConfig"`
}

// 构建环境持久化配置
type CiEnvPersistenceConfig []*CiEnvPersistenceConfigItem

type CiEnvPersistenceConfigItem struct {
	PvcName   string `json:"pvcName"`
	MountPath string `json:"mountPath"`
	SubPath   string `json:"subPath"`
}
