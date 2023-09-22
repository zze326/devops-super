package user

import (
	"context"
	"devops-super/api/user/v1"
	"devops-super/internal/model/do"
	"devops-super/internal/service"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = service.User().Del(ctx, &do.User{Id: req.Id})
	return
}
