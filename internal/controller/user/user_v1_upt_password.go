package user

import (
	"context"
	"devops-super/api/user/v1"
	"devops-super/internal/model/do"
	"devops-super/internal/service"
)

func (c *ControllerV1) UptPassword(ctx context.Context, req *v1.UptPasswordReq) (res *v1.UptPasswordRes, err error) {
	err = service.User().Upt(ctx, &do.User{Id: req.Id, Password: req.Password})
	return
}
