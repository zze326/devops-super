// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	model0 "devops-super/internal/model"
)

type (
	IContext interface {
		Init(ctx context.Context) error
		RefreshCasbin(ctx context.Context) error
		Ctx() *model0.ServiceContext
	}
)

var (
	localContext IContext
)

func Context() IContext {
	if localContext == nil {
		panic("implement not found for interface IContext, forgot register?")
	}
	return localContext
}

func RegisterContext(i IContext) {
	localContext = i
}
