package host

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/host/v1"
)

func (c *ControllerV1) TestSsh(ctx context.Context, req *v1.TestSshReq) (res *v1.TestSshRes, err error) {
	err = service.Host().TestSSH(ctx, req.Id)
	return
}
