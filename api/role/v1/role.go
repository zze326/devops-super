package v1

import (
	"devops-super/api"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/frame/g"
)

type GetPageLstReq struct {
	g.Meta `method:"get" path:"/role/page-list" summary:"分页获取角色列表" tags:"角色"`
	*api.PageLstReq
}

type GetPageLstRes struct {
	*api.PageLstRes[*entity.Role]
}

type AddReq struct {
	g.Meta `method:"post" path:"/role" summary:"添加角色" tags:"角色"`
	*mid.Role
}

type AddRes struct{}

type UptReq struct {
	g.Meta `method:"put" path:"/role/{id}" summary:"更新角色" tags:"角色"`
	Id     int ` v:"min:1#id必须" path:"id"`
	*mid.Role
}

type UptRes struct{}

type DelReq struct {
	g.Meta `method:"delete" path:"/role/{id}" summary:"删除角色" tags:"角色"`
	Id     int ` v:"min:1#id必须" path:"id"`
}

type DelRes struct{}

type GetLstReq struct {
	g.Meta `method:"get" path:"/role/list" summary:"获取所有角色列表" tags:"角色"`
	*api.PageLstReq
}

type GetLstRes struct {
	List []*entity.Role `json:"list"`
}
