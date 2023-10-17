package ci_pipeline_run

import (
	"context"
	"devops-super/api"
	"devops-super/internal/dao"
	"devops-super/internal/model/entity"
	"devops-super/internal/service"
)

type sCiPipelineRun struct{}

var cols = dao.CiPipelineRun.Columns()

func init() {
	service.RegisterCiPipelineRun(New())
}

func New() *sCiPipelineRun {
	return &sCiPipelineRun{}
}

func (*sCiPipelineRun) GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.CiPipelineRun], err error) {
	out = &api.PageLstRes[*entity.CiPipelineRun]{}
	m := dao.CiPipelineRun.Ctx(ctx).Safe(true)
	if pipelineId := in.Wheres.Get("pipelineId"); !pipelineId.IsNil() {
		m = m.Where(m.Builder().Where(cols.PipelineId, pipelineId.Int()))
	}
	err = m.Offset(in.Offset()).Limit(in.Limit()).OrderDesc(cols.Id).
		ScanAndCount(&out.List, &out.Total, false)
	return
}
