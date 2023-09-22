package permission

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/util/gconv"

	"devops-super/api/permission/v1"
)

func (c *ControllerV1) Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error) {
	uptDo := new(do.Permission)
	if err = gconv.Struct(req, uptDo); err != nil {
		return
	}
	err = service.Permission().Upt(ctx, uptDo)
	return
}
