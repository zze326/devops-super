// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// KubernetesConfig is the golang structure for table kubernetes_config.
type KubernetesConfig struct {
	Id        int         `json:"id"        description:""`     //
	Name      string      `json:"name"      description:"名称"`   // 名称
	Config    string      `json:"config"    description:"配置内容"` // 配置内容
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"` // 更新时间
}
