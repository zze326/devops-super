// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CiPipelineRun is the golang structure for table ci_pipeline_run.
type CiPipelineRun struct {
	Id         int         `json:"id"         description:""`                        //
	PipelineId int         `json:"pipelineId" description:"流水线 id"`                  // 流水线 id
	PodName    string      `json:"podName"    description:"Pod 名称"`                  // Pod 名称
	Namespace  string      `json:"namespace"  description:"名称空间"`                    // 名称空间
	Status     int         `json:"status"     description:"状态:0-运行中,1:成功,2:失败,3:取消"` // 状态:0-运行中,1:成功,2:失败,3:取消
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"更新时间"`                    // 更新时间
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`                    // 创建时间
}
