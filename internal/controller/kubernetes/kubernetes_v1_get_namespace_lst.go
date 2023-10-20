package kubernetes

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"

	"devops-super/api/kubernetes/v1"
)

func (c *ControllerV1) GetNamespaceLst(ctx context.Context, req *v1.GetNamespaceLstReq) (res *v1.GetNamespaceLstRes, err error) {
	res = new(v1.GetNamespaceLstRes)
	config, err := service.Secret().GetKubernetesConfig(ctx, &do.Secret{Id: req.SecretId})
	if err != nil {
		return
	}

	res.Namespaces, err = service.Kubernetes().GetNamespaces(ctx, config.Text)
	return
}
