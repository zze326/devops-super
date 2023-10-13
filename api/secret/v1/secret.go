package v1

import (
	"devops-super/api"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/frame/g"
)

type GetPageLstReq struct {
	g.Meta `method:"get" path:"/secret/page-list" summary:"分页获取秘钥列表" tags:"秘钥"`
	*api.PageLstReq
}

type GetPageLstRes struct {
	*api.PageLstRes[*entity.Secret]
}

type AddReq struct {
	g.Meta `method:"post" path:"/secret" summary:"添加秘钥" tags:"秘钥"`
	*mid.Secret
}

type AddRes struct{}

type UptReq struct {
	g.Meta `method:"put" path:"/secret/{id}" summary:"更新秘钥" tags:"秘钥"`
	Id     int ` v:"min:1#id必须" path:"id"`
	*mid.Secret
}

type UptRes struct{}

type DelReq struct {
	g.Meta `method:"delete" path:"/secret/{id}" summary:"删除秘钥" tags:"秘钥"`
	Id     int ` v:"min:1#id必须" path:"id"`
}

type DelRes struct{}

type GetLstReq struct {
	g.Meta `method:"get" path:"/secret/list" summary:"获取所有秘钥列表" tags:"秘钥"`
	Type   int `p:"type"`
}

type GetLstRes struct {
	List []*entity.Secret `json:"list"`
}
