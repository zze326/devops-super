// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// CiPipeline is the golang structure for table ci_pipeline.
type CiPipeline struct {
	Id                  int         `json:"id"                  description:""`                         //
	Name                string      `json:"name"                description:"名称"`                       // 名称
	KubernetesConfigId  int         `json:"kubernetesConfigId"  description:"关联的 Kubernetes Config id"` // 关联的 Kubernetes Config id
	KubernetesNamespace string      `json:"kubernetesNamespace" description:"Pod 所在命名空间"`               // Pod 所在命名空间
	Params              *gjson.Json `json:"params"              description:"构建参数"`                     // 构建参数
	Config              *gjson.Json `json:"config"              description:"配置"`                       // 配置
	Desc                string      `json:"desc"                description:"描述"`                       // 描述
	UpdatedAt           *gtime.Time `json:"updatedAt"           description:"更新时间"`                     // 更新时间
}
