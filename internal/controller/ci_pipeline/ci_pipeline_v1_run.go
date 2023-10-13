package ci_pipeline

import (
	"context"
	"devops-super/api/ci_pipeline/v1"
	"devops-super/internal/service"
)

func (c *ControllerV1) Run(ctx context.Context, req *v1.RunReq) (res *v1.RunRes, err error) {
	err = service.CiPipeline().Run(ctx, req.Id)
	return
}
