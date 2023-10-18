// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"devops-super/api"
	"devops-super/internal/model/entity"
)

type (
	ICiPipelineRun interface {
		GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.CiPipelineRun], err error)
		WsLog(ctx context.Context, id int) (err error)
		WsPageLst(ctx context.Context) error
	}
)

var (
	localCiPipelineRun ICiPipelineRun
)

func CiPipelineRun() ICiPipelineRun {
	if localCiPipelineRun == nil {
		panic("implement not found for interface ICiPipelineRun, forgot register?")
	}
	return localCiPipelineRun
}

func RegisterCiPipelineRun(i ICiPipelineRun) {
	localCiPipelineRun = i
}
