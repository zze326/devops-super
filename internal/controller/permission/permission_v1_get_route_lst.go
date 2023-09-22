package permission

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/permission/v1"
)

func (c *ControllerV1) GetRouteLst(ctx context.Context, req *v1.GetRouteLstReq) (res *v1.GetRouteLstRes, err error) {
	res = new(v1.GetRouteLstRes)
	res.List, err = service.Permission().GetRouteLst(ctx)
	return
}
