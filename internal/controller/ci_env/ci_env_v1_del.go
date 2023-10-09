package ci_env

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"

	"devops-super/api/ci_env/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = service.CiEnv().Del(ctx, &do.CiEnv{Id: req.Id})
	return
}
