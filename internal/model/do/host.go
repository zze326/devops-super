// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Host is the golang structure of table host for DAO operations like Where/Data.
type Host struct {
	g.Meta      `orm:"table:host, do:true"`
	Id          interface{} //
	Name        interface{} // 名称
	Host        interface{} // 主机名或IP
	Port        interface{} // 端口
	Username    interface{} // 用户名
	Password    interface{} // 密码
	PrivateKey  interface{} // 私钥
	UseKey      interface{} // 是否使用公钥连接
	Desc        interface{} // 描述
	SaveSession interface{} // 是否保存会话
	UpdatedAt   *gtime.Time // 更新时间
}
