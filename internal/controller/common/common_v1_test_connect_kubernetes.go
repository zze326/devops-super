package common

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/common/v1"
)

func (c *ControllerV1) TestConnectKubernetes(ctx context.Context, req *v1.TestConnectKubernetesReq) (res *v1.TestConnectKubernetesRes, err error) {
	err = service.Common().TestConnectKubernetes(ctx, req.Config)
	return
}
