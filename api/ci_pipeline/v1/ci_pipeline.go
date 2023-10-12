package v1

import (
	"devops-super/api"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type GetPageLstReq struct {
	g.Meta `method:"get" path:"/ci-pipeline/page-list" summary:"分页获取 CI 流水线列表" tags:"CI 流水线"`
	*api.PageLstReq
}

type GetPageLstRes struct {
	*api.PageLstRes[*entity.CiPipeline]
}

type AddReq struct {
	g.Meta `method:"post" path:"/ci-pipeline" summary:"添加 CI 流水线" tags:"CI 流水线"`
	*mid.CiPipeline
}

type AddRes struct{}

type UptReq struct {
	g.Meta `method:"put" path:"/ci-pipeline/{id}" summary:"更新 CI 流水线" tags:"CI 流水线"`
	Id     int ` v:"min:1#id必须" path:"id"`
	*mid.CiPipeline
}

type UptRes struct{}

type UptConfigReq struct {
	g.Meta `method:"patch" path:"/ci-pipeline/{id}/config" summary:"更新 CI 流水线配置" tags:"CI 流水线"`
	Id     int         `v:"min:1#id必须" path:"id"`
	Config *gjson.Json `v:"required" json:"config"`
}

type UptConfigRes struct{}

type DelReq struct {
	g.Meta `method:"delete" path:"/ci-pipeline/{id}" summary:"删除 CI 流水线" tags:"CI 流水线"`
	Id     int ` v:"min:1#id必须" path:"id"`
}

type DelRes struct{}

type GetLstReq struct {
	g.Meta `method:"get" path:"/ci-pipeline/list" summary:"获取所有 CI 流水线列表" tags:"CI 流水线"`
	*api.PageLstReq
}

type GetLstRes struct {
	List []*entity.CiPipeline `json:"list"`
}

type GetConfigReq struct {
	g.Meta `method:"get" path:"/ci-pipeline/{id}/config" summary:"获取 CI 流水线配置" tags:"CI 流水线"`
	Id     int ` v:"min:1#id必须" path:"id"`
}

type GetConfigRes struct {
	Config *gjson.Json    `json:"config"`
	EnvMap map[int]string `json:"envMap"`
}
