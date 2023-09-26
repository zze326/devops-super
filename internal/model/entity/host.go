// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Host is the golang structure for table host.
type Host struct {
	Id          int         `json:"id"          description:""`         //
	Name        string      `json:"name"        description:"名称"`       // 名称
	HostAddr    string      `json:"hostAddr"    description:"主机名或IP"`   // 主机名或IP
	Port        int64       `json:"port"        description:"端口"`       // 端口
	Username    string      `json:"username"    description:"用户名"`      // 用户名
	Password    string      `json:"password"    description:"密码"`       // 密码
	PrivateKey  string      `json:"privateKey"  description:"私钥"`       // 私钥
	UseKey      bool        `json:"useKey"      description:"是否使用公钥连接"` // 是否使用公钥连接
	Desc        string      `json:"desc"        description:"描述"`       // 描述
	SaveSession bool        `json:"saveSession" description:"是否保存会话"`   // 是否保存会话
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`     // 更新时间
	HostGroupId int         `json:"hostGroupId" description:"主机组 id"`   // 主机组 id
}
