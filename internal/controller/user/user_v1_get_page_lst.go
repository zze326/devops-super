package user

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/user/v1"
)

func (c *ControllerV1) GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error) {
	res = new(v1.GetPageLstRes)
	res.PageLstRes, err = service.User().GetPageLst(ctx, req.PageLstReq)
	return
}
