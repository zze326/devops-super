package kubernetes_config

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/kubernetes_config/v1"
)

func (c *ControllerV1) GetLst(ctx context.Context, req *v1.GetLstReq) (res *v1.GetLstRes, err error) {
	res = new(v1.GetLstRes)
	res.List, err = service.KubernetesConfig().GetLst(ctx)
	return
}
