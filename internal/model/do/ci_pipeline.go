// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CiPipeline is the golang structure of table ci_pipeline for DAO operations like Where/Data.
type CiPipeline struct {
	g.Meta             `orm:"table:ci_pipeline, do:true"`
	Id                 interface{} //
	Name               interface{} // 名称
	KubernetesConfigId interface{} // 关联的 Kubernetes Config id
	Config             *gjson.Json // 配置
	Desc               interface{} // 描述
	UpdatedAt          *gtime.Time // 更新时间
}
