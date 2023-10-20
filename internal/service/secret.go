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
	"devops-super/internal/model/mid"
)

type (
	ISecret interface {
		Add(ctx context.Context, in *entity.Secret) (err error)
		Upt(ctx context.Context, in *do.Secret) (err error)
		GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.Secret], err error)
		GetLst(ctx context.Context, in *do.Secret) (out []*entity.Secret, err error)
		Get(ctx context.Context, in *do.Secret) (out *entity.Secret, err error)
		Del(ctx context.Context, in *do.Secret) (err error)
		GetKubernetesConfig(ctx context.Context, in *do.Secret) (*mid.TextContent, error)
	}
)

var (
	localSecret ISecret
)

func Secret() ISecret {
	if localSecret == nil {
		panic("implement not found for interface ISecret, forgot register?")
	}
	return localSecret
}

func RegisterSecret(i ISecret) {
	localSecret = i
}
