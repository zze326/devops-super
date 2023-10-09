package kubernetes_config

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"

	"devops-super/api/kubernetes_config/v1"
)

func (c *ControllerV1) Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error) {
	err = service.KubernetesConfig().Del(ctx, &do.KubernetesConfig{Id: req.Id})
	return
}
