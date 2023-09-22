package role

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/role/v1"
)

func (c *ControllerV1) GetLst(ctx context.Context, req *v1.GetLstReq) (res *v1.GetLstRes, err error) {
	res = new(v1.GetLstRes)
	res.List, err = service.Role().GetLst(ctx)
	return
}
