package ci_pipeline

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/util/gconv"

	"devops-super/api/ci_pipeline/v1"
)

func (c *ControllerV1) Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error) {
	in := new(do.CiPipeline)
	if err = gconv.Struct(req, in); err != nil {
		return
	}
	err = service.CiPipeline().Upt(ctx, in)
	return
}
