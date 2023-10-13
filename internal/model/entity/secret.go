// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// Secret is the golang structure for table secret.
type Secret struct {
	Id        int         `json:"id"        description:""`                               //
	Name      string      `json:"name"      description:"名称"`                             // 名称
	Type      int         `json:"type"      description:"类型:1-git认证,2-Kubernetes config"` // 类型:1-git认证,2-Kubernetes config
	Content   *gjson.Json `json:"content"   description:"认证配置内容"`                         // 认证配置内容
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`                           // 更新时间
}
