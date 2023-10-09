// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CiEnv is the golang structure for table ci_env.
type CiEnv struct {
	Id         int         `json:"id"         description:""`                            //
	Name       string      `json:"name"       description:"环境名称"`                        // 环境名称
	Image      string      `json:"image"      description:"镜像"`                          // 镜像
	SecretName string      `json:"secretName" description:"Kubernetes Secret 名称，拉取镜像使用"` // Kubernetes Secret 名称，拉取镜像使用
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"更新时间"`                        // 更新时间
}
