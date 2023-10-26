// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CiPipelineRun is the golang structure of table ci_pipeline_run for DAO operations like Where/Data.
type CiPipelineRun struct {
	g.Meta     `orm:"table:ci_pipeline_run, do:true"`
	Id         interface{} //
	PipelineId interface{} // 流水线 id
	PodName    interface{} // Pod 名称
	Namespace  interface{} // 名称空间
	Status     interface{} // 状态:0-运行中,1:成功,2:失败,3:取消
	UpdatedAt  *gtime.Time // 更新时间
	CreatedAt  *gtime.Time // 创建时间
}
