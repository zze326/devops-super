package public

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/public/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, _ *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
