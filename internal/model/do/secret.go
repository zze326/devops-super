// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Secret is the golang structure of table secret for DAO operations like Where/Data.
type Secret struct {
	g.Meta    `orm:"table:secret, do:true"`
	Id        interface{} //
	Name      interface{} // 名称
	Type      interface{} // 类型:1-git认证,2-Kubernetes config
	Content   *gjson.Json // 认证配置内容
	UpdatedAt *gtime.Time // 更新时间
}
