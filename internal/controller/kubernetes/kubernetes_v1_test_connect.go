package kubernetes

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/kubernetes/v1"
)

func (c *ControllerV1) TestConnect(ctx context.Context, req *v1.TestConnectReq) (res *v1.TestConnectRes, err error) {
	err = service.Kubernetes().TestConnect(ctx, req.Config)
	return
}
