package v1

import (
	"devops-super/api"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/frame/g"
)

type GetPageLstReq struct {
	g.Meta `method:"get" path:"/ci-env/page-list" summary:"分页获取 CI 环境列表" tags:"CI 环境"`
	*api.PageLstReq
}

type GetPageLstRes struct {
	*api.PageLstRes[*entity.CiEnv]
}

type AddReq struct {
	g.Meta `method:"post" path:"/ci-env" summary:"添加 CI 环境" tags:"CI 环境"`
	*mid.CiEnv
}

type AddRes struct{}

type UptReq struct {
	g.Meta `method:"put" path:"/ci-env/{id}" summary:"更新 CI 环境" tags:"CI 环境"`
	Id     int ` v:"min:1#id必须" path:"id"`
	*mid.CiEnv
}

type UptRes struct{}

type DelReq struct {
	g.Meta `method:"delete" path:"/ci-env/{id}" summary:"删除 CI 环境" tags:"CI 环境"`
	Id     int ` v:"min:1#id必须" path:"id"`
}

type DelRes struct{}

type GetLstReq struct {
	g.Meta `method:"get" path:"/ci-env/list" summary:"获取所有 CI 环境列表" tags:"CI 环境"`
	*api.PageLstReq
}

type GetLstRes struct {
	List []*entity.CiEnv `json:"list"`
}
