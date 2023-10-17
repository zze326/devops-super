package ci_pipeline_run

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/ci_pipeline_run/v1"
)

func (c *ControllerV1) WsLog(ctx context.Context, req *v1.WsLogReq) (res *v1.WsLogRes, err error) {
	err = service.CiPipelineRun().WsLog(ctx, req.Id)
	return
}
