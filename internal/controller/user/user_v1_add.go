package user

import (
	"context"
	"devops-super/internal/model/entity"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/util/gconv"

	"devops-super/api/user/v1"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	eUser := new(entity.User)
	if err = gconv.Struct(req, eUser); err != nil {
		return
	}

	eUser.Password = "devops.zze"
	err = service.User().Add(ctx, eUser)
	return
}
