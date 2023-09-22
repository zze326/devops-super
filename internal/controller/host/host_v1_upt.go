package host

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/util/gconv"

	"devops-super/api/host/v1"
)

func (c *ControllerV1) Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error) {
	in := new(do.Host)
	if err = gconv.Struct(req, in); err != nil {
		return
	}
	err = service.Host().Upt(ctx, in)
	return
}
