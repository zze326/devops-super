// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

// Permission is the golang structure of table permission for DAO operations like Where/Data.
type Permission struct {
	g.Meta     `orm:"table:permission, do:true"`
	Id         interface{} //
	Title      interface{} // 标题
	Name       interface{} // 路由名称
	Type       interface{} // 类型:1-目录,2-菜单,3-功能
	FRoute     interface{} // 前端路由路径
	BRoutes    *gjson.Json // 后端路由路径
	Redirect   interface{} // 重定向路径
	Icon       interface{} // 图标
	Rank       interface{} // 排序
	ShowLink   interface{} // 是否在菜单中展示
	ShowParent interface{} // 是否展示父级菜单
	KeepAlive  interface{} // 页面缓存
	ParentId   interface{} // 父级权限 id
}
