package secret

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"

	"devops-super/api/secret/v1"
)

func (c *ControllerV1) GetLst(ctx context.Context, req *v1.GetLstReq) (res *v1.GetLstRes, err error) {
	res = new(v1.GetLstRes)
	in := new(do.Secret)
	if req.Type != 0 {
		in.Type = req.Type
	}
	res.List, err = service.Secret().GetLst(ctx, in)
	return
}
