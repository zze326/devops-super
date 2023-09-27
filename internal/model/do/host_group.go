// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HostGroup is the golang structure of table host_group for DAO operations like Where/Data.
type HostGroup struct {
	g.Meta    `orm:"table:host_group, do:true"`
	Id        interface{} //
	Name      interface{} // 主机组名称
	Rank      interface{} // 排序
	ParentId  interface{} // 上级主机组 id
	UpdatedAt *gtime.Time // 更新时间
	RoleIds   *gjson.Json // 可访问的角色 id 列表
	UserIds   *gjson.Json // 可访问的用户 id 列表
}
