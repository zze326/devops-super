package ci_pipeline

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/ci_pipeline/v1"
)

func (c *ControllerV1) Clone(ctx context.Context, req *v1.CloneReq) (res *v1.CloneRes, err error) {
	err = service.CiPipeline().Clone(ctx, req.Id, req.NewName)
	return
}
