package v1

import (
	"devops-super/api"
	"devops-super/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetPageLstReq struct {
	g.Meta `method:"get" path:"/ci-pipeline-run/page-list" summary:"分页获取 CI 流水线运行记录列表" tags:"CI 流水线运行记录"`
	*api.PageLstReq
}

type GetPageLstRes struct {
	*api.PageLstRes[*entity.CiPipelineRun]
}
