// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
)

// Permission is the golang structure for table permission.
type Permission struct {
	Id         int         `json:"id"         description:""`                  //
	Title      string      `json:"title"      description:"标题"`                // 标题
	Name       string      `json:"name"       description:"路由名称"`              // 路由名称
	Type       int         `json:"type"       description:"类型:1-目录,2-菜单,3-功能"` // 类型:1-目录,2-菜单,3-功能
	FRoute     string      `json:"fRoute"     description:"前端路由路径"`            // 前端路由路径
	BRoutes    *gjson.Json `json:"bRoutes"    description:"后端路由路径"`            // 后端路由路径
	Redirect   string      `json:"redirect"   description:"重定向路径"`             // 重定向路径
	Icon       string      `json:"icon"       description:"图标"`                // 图标
	Rank       int         `json:"rank"       description:"排序"`                // 排序
	ShowLink   bool        `json:"showLink"   description:"是否在菜单中展示"`          // 是否在菜单中展示
	ShowParent bool        `json:"showParent" description:"是否展示父级菜单"`          // 是否展示父级菜单
	KeepAlive  bool        `json:"keepAlive"  description:"页面缓存"`              // 页面缓存
	ParentId   int         `json:"parentId"   description:"父级权限 id"`           // 父级权限 id
}
