// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Dept is the golang structure for table dept.
type Dept struct {
	Id        int         `json:"id"        description:""`        //
	Name      string      `json:"name"      description:"部门名称"`    // 部门名称
	Rank      int         `json:"rank"      description:"排序"`      // 排序
	ParentId  int         `json:"parentId"  description:"上级部门 id"` // 上级部门 id
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`    // 更新时间
}
