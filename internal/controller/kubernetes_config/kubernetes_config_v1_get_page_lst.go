package kubernetes_config

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/kubernetes_config/v1"
)

func (c *ControllerV1) GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error) {
	res = new(v1.GetPageLstRes)
	res.PageLstRes, err = service.KubernetesConfig().GetPageLst(ctx, req.PageLstReq)
	return
}
