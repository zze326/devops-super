// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"devops-super/api"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
)

type (
	IKubernetesConfig interface {
		Add(ctx context.Context, in *entity.KubernetesConfig) (err error)
		Upt(ctx context.Context, in *do.KubernetesConfig) (err error)
		GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.KubernetesConfig], err error)
		GetLst(ctx context.Context) (out []*entity.KubernetesConfig, err error)
		Get(ctx context.Context, in *do.KubernetesConfig) (out *entity.KubernetesConfig, err error)
		Del(ctx context.Context, in *do.KubernetesConfig) (err error)
	}
)

var (
	localKubernetesConfig IKubernetesConfig
)

func KubernetesConfig() IKubernetesConfig {
	if localKubernetesConfig == nil {
		panic("implement not found for interface IKubernetesConfig, forgot register?")
	}
	return localKubernetesConfig
}

func RegisterKubernetesConfig(i IKubernetesConfig) {
	localKubernetesConfig = i
}
