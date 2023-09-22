// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
)

type (
	IPermission interface {
		Add(ctx context.Context, in *entity.Permission) (err error)
		Get(ctx context.Context, in *do.Permission) (out *entity.Permission, err error)
		GetLst(ctx context.Context, search string) (out []*entity.Permission, err error)
		Upt(ctx context.Context, in *do.Permission) (err error)
		Del(ctx context.Context, in *do.Permission) (err error)
		SystemRequired(ctx context.Context) (ePermission *entity.Permission, err error)
		GetRouteLst(ctx context.Context) (out []*mid.Route, err error)
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
