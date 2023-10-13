package secret

import (
	"context"
	"devops-super/internal/model/entity"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/util/gconv"

	"devops-super/api/secret/v1"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	in := new(entity.Secret)
	if err = gconv.Struct(req, in); err != nil {
		return
	}
	err = service.Secret().Add(ctx, in)
	return
}
