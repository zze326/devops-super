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
	ICiEnv interface {
		Add(ctx context.Context, in *entity.CiEnv) (err error)
		Upt(ctx context.Context, in *do.CiEnv) (err error)
		GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.CiEnv], err error)
		GetLst(ctx context.Context) (out []*entity.CiEnv, err error)
		Get(ctx context.Context, in *do.CiEnv) (out *entity.CiEnv, err error)
		Del(ctx context.Context, in *do.CiEnv) (err error)
	}
)

var (
	localCiEnv ICiEnv
)

func CiEnv() ICiEnv {
	if localCiEnv == nil {
		panic("implement not found for interface ICiEnv, forgot register?")
	}
	return localCiEnv
}

func RegisterCiEnv(i ICiEnv) {
	localCiEnv = i
}
