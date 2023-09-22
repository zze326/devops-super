package role

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"

	"devops-super/api/role/v1"
)

func (c *ControllerV1) UptPermission(ctx context.Context, req *v1.UptPermissionReq) (res *v1.UptPermissionRes, err error) {
	err = service.Role().Upt(ctx, &do.Role{Id: req.Id, Permission: req.PermissionIds})
	return
}
