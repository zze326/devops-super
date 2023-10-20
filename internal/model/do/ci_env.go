// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CiEnv is the golang structure of table ci_env for DAO operations like Where/Data.
type CiEnv struct {
	g.Meta            `orm:"table:ci_env, do:true"`
	Id                interface{} //
	Name              interface{} // 环境名称
	Image             interface{} // 镜像
	SecretName        interface{} // Kubernetes Secret 名称，拉取镜像使用
	PersistenceConfig *gjson.Json // 持久化配置
	UpdatedAt         *gtime.Time // 更新时间
}
