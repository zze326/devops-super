package kubernetes_config

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/kubernetes_config/v1"
)

func (c *ControllerV1) GetPartialLst(ctx context.Context, req *v1.GetPartialLstReq) (res *v1.GetPartialLstRes, err error) {
	res = new(v1.GetPartialLstRes)
	res.List, err = service.KubernetesConfig().GetPartialLst(ctx)
	return
}
