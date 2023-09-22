package dept

import (
	"context"
	"devops-super/internal/model/entity"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/util/gconv"

	"devops-super/api/dept/v1"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	in := new(entity.Dept)
	if err = gconv.Struct(req, in); err != nil {
		return
	}

	err = service.Dept().Add(ctx, in)
	return
}
