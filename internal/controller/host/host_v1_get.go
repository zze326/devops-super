package host

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"

	"devops-super/api/host/v1"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	res = new(v1.GetRes)
	res.Host, err = service.Host().Get(ctx, &do.Host{Id: req.Id})
	return
}
