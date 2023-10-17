package ci_pipeline_run

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/ci_pipeline_run/v1"
)

func (c *ControllerV1) GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error) {
	res = new(v1.GetPageLstRes)
	res.PageLstRes, err = service.CiPipelineRun().GetPageLst(ctx, req.PageLstReq)
	return
}
