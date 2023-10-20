package kubernetes

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"

	"devops-super/api/kubernetes/v1"
)

func (c *ControllerV1) GetPersistentVolumeClaimLst(ctx context.Context, req *v1.GetPersistentVolumeClaimLstReq) (res *v1.GetPersistentVolumeClaimLstRes, err error) {
	res = new(v1.GetPersistentVolumeClaimLstRes)
	config, err := service.Secret().GetKubernetesConfig(ctx, &do.Secret{Id: req.SecretId})
	if err != nil {
		return
	}

	res.Pvcs, err = service.Kubernetes().GetPersistentVolumeClaims(ctx, config.Text, req.Namespace)
	return
}
