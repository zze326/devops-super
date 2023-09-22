package permission

import (
	"context"
	"devops-super/internal/model/entity"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/util/gconv"

	"devops-super/api/permission/v1"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	ePermission := new(entity.Permission)
	if err = gconv.Struct(req, ePermission); err != nil {
		return
	}
	err = service.Permission().Add(ctx, ePermission)
	return
}
