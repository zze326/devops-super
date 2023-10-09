package ci_env

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/util/gconv"

	"devops-super/api/ci_env/v1"
)

func (c *ControllerV1) Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error) {
	in := new(do.CiEnv)
	if err = gconv.Struct(req, in); err != nil {
		return
	}
	err = service.CiEnv().Upt(ctx, in)
	return
}
