// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        interface{} //
	Username  interface{} // 用户名
	Password  interface{} // 密码
	Phone     interface{} // 手机号码
	Email     interface{} // 邮箱
	RealName  interface{} // 真实姓名
	Enabled   interface{} // 是否启用状态
	RoleIds   *gjson.Json // 角色 id
	DeptId    interface{} // 所属部门 id
	UpdatedAt *gtime.Time // 更新时间
}
