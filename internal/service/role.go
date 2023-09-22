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
	"devops-super/internal/model/entity/comb"
)

type (
	IRole interface {
		Add(ctx context.Context, in *entity.Role) (err error)
		Upt(ctx context.Context, in *do.Role) (err error)
		GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.Role], err error)
		GetLst(ctx context.Context) (out []*entity.Role, err error)
		GetCombList(ctx context.Context) (out []*comb.Role, err error)
		Get(ctx context.Context, in *do.Role) (out *entity.Role, err error)
		Del(ctx context.Context, in *do.Role) (err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
