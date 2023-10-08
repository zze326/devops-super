package host

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/host/v1"
)

func (c *ControllerV1) GetAuthorizedLst(ctx context.Context, req *v1.GetAuthorizedLstReq) (res *v1.GetAuthorizedLstRes, err error) {
	res = new(v1.GetAuthorizedLstRes)
	res.List, err = service.Host().GetAuthorizedLst(ctx)
	return
}
