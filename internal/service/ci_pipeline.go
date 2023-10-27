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

	"github.com/gogf/gf/v2/encoding/gjson"
)

type (
	ICiPipeline interface {
		Add(ctx context.Context, in *entity.CiPipeline) (err error)
		Upt(ctx context.Context, in *do.CiPipeline) (err error)
		GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.CiPipeline], err error)
		GetLst(ctx context.Context) (out []*entity.CiPipeline, err error)
		Get(ctx context.Context, in *do.CiPipeline) (out *entity.CiPipeline, err error)
		Del(ctx context.Context, in *do.CiPipeline) (err error)
		Run(ctx context.Context, id int, params *gjson.Json) (err error)
	}
)

var (
	localCiPipeline ICiPipeline
)

func CiPipeline() ICiPipeline {
	if localCiPipeline == nil {
		panic("implement not found for interface ICiPipeline, forgot register?")
	}
	return localCiPipeline
}

func RegisterCiPipeline(i ICiPipeline) {
	localCiPipeline = i
}
