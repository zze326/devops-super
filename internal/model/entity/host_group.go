// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HostGroup is the golang structure for table host_group.
type HostGroup struct {
	Id        int         `json:"id"        description:""`         //
	Name      string      `json:"name"      description:"主机组名称"`    // 主机组名称
	Rank      int         `json:"rank"      description:"排序"`       // 排序
	ParentId  int         `json:"parentId"  description:"上级主机组 id"` // 上级主机组 id
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`     // 更新时间
}
