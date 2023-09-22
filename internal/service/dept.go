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
	IDept interface {
		Add(ctx context.Context, in *entity.Dept) (err error)
		Get(ctx context.Context, in *do.Dept) (out *entity.Dept, err error)
		GetLst(ctx context.Context, search string) (out []*entity.Dept, err error)
		Upt(ctx context.Context, in *do.Dept) (err error)
		Del(ctx context.Context, in *do.Dept) (err error)
	}
)

var (
	localDept IDept
)

func Dept() IDept {
	if localDept == nil {
		panic("implement not found for interface IDept, forgot register?")
	}
	return localDept
}

func RegisterDept(i IDept) {
	localDept = i
}
