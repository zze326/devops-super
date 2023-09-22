package dept

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/util/gconv"

	"devops-super/api/dept/v1"
)

func (c *ControllerV1) Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error) {
	in := new(do.Dept)
	if err = gconv.Struct(req, in); err != nil {
		return
	}
	err = service.Dept().Upt(ctx, in)
	return
}
