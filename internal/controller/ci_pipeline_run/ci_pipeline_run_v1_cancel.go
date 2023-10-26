package ci_pipeline_run

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/ci_pipeline_run/v1"
)

func (c *ControllerV1) Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error) {
	err = service.CiPipelineRun().Cancel(ctx, req.Id)
	return
}
