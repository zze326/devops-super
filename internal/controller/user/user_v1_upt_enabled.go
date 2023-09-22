package user

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"

	"devops-super/api/user/v1"
)

func (c *ControllerV1) UptEnabled(ctx context.Context, req *v1.UptEnabledReq) (res *v1.UptEnabledRes, err error) {
	err = service.User().Upt(ctx, &do.User{Id: req.Id, Enabled: req.Enabled})
	return
}
