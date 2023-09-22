package v1

import (
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/frame/g"
)

type AddReq struct {
	g.Meta `path:"/permission" method:"post" tags:"权限" summary:"创建权限"`
	*mid.Permission
}

type AddRes struct{}

type GetLstReq struct {
	g.Meta `path:"/permission/list" method:"get" tags:"权限" summary:"获取权限列表"`
	Search string `p:"search"`
}

type GetLstRes struct {
	List []*entity.Permission `json:"list"`
}

type UptReq struct {
	g.Meta `path:"/permission/{id}" method:"put" tags:"权限" summary:"更新权限"`
	Id     int ` v:"min:1#id必须" path:"id"`
	*mid.Permission
}

type UptRes struct{}

type UptShowLinkReq struct {
	g.Meta  `path:"/permission/{id}/show-link" method:"patch" tags:"权限" summary:"更新是否显示在菜单"`
	Id      int  `v:"min:1#id必须" path:"id" `
	Enabled bool `v:"required" json:"enabled"`
}

type UptShowLinkRes struct{}

type DelReq struct {
	g.Meta `path:"/permission/{id}" method:"delete" tags:"权限" summary:"删除权限"`
	Id     int `v:"min:1#id必须" path:"id" `
}

type DelRes struct{}

type GetRouteLstReq struct {
	g.Meta `path:"/permission/route-list" method:"get" tags:"权限" summary:"获取前端路由列表"`
}

type GetRouteLstRes struct {
	List []*mid.Route `json:"list"`
}
