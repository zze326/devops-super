package public

import (
	"context"

	"devops-super/api/public/v1"
)

func (c *ControllerV1) Ping(ctx context.Context, req *v1.PingReq) (res *v1.PingRes, err error) {
	return
}
