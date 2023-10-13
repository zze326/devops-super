package secret

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"

	"devops-super/api/secret/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = service.Secret().Del(ctx, &do.Secret{Id: req.Id})
	return
}
