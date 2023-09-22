// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta     `orm:"table:role, do:true"`
	Id         interface{} //
	Name       interface{} // 角色名称
	Code       interface{} // 角色代码
	Permission *gjson.Json // 关联权限
	UpdatedAt  *gtime.Time // 更新时间
}
