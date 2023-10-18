package ci_pipeline_run

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/ci_pipeline_run/v1"
)

func (c *ControllerV1) WsPageLst(ctx context.Context, req *v1.WsPageLstReq) (res *v1.WsPageLstRes, err error) {
	err = service.CiPipelineRun().WsPageLst(ctx)
	return
}
