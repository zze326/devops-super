// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Dept is the golang structure of table dept for DAO operations like Where/Data.
type Dept struct {
	g.Meta    `orm:"table:dept, do:true"`
	Id        interface{} //
	Name      interface{} // 部门名称
	Rank      interface{} // 排序
	ParentId  interface{} // 上级部门 id
	UpdatedAt *gtime.Time // 更新时间
}
