package mid

import "github.com/gogf/gf/v2/encoding/gjson"

type User struct {
	Username string      `v:"required|length:4,30" json:"username"`
	Phone    string      `json:"phone"`
	Email    string      `json:"email"`
	RealName string      `v:"required" json:"realName"`
	RoleIds  *gjson.Json `json:"roleIds"`
	DeptId   int         `v:"required" json:"deptId"`
}
