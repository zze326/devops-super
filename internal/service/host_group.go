// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
)

type (
	IHostGroup interface {
		Add(ctx context.Context, in *entity.HostGroup) (err error)
		Get(ctx context.Context, in *do.HostGroup) (out *entity.HostGroup, err error)
		GetLst(ctx context.Context, search string) (out []*entity.HostGroup, err error)
		Upt(ctx context.Context, in *do.HostGroup) (err error)
		Del(ctx context.Context, in *do.HostGroup) (err error)
	}
)

var (
	localHostGroup IHostGroup
)

func HostGroup() IHostGroup {
	if localHostGroup == nil {
		panic("implement not found for interface IHostGroup, forgot register?")
	}
	return localHostGroup
}

func RegisterHostGroup(i IHostGroup) {
	localHostGroup = i
}
