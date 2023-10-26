package ci_pipeline

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"

	"devops-super/api/ci_pipeline/v1"
)

func (c *ControllerV1) UptConfig(ctx context.Context, req *v1.UptConfigReq) (res *v1.UptConfigRes, err error) {
	err = service.CiPipeline().Upt(ctx, &do.CiPipeline{Id: req.Id, Config: req.Config, Params: req.Params})
	return
}
