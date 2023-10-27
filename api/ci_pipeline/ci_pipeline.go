// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ci_pipeline

import (
	"context"

	"devops-super/api/ci_pipeline/v1"
)

type ICiPipelineV1 interface {
	GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error)
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error)
	UptConfig(ctx context.Context, req *v1.UptConfigReq) (res *v1.UptConfigRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
	GetLst(ctx context.Context, req *v1.GetLstReq) (res *v1.GetLstRes, err error)
	GetConfig(ctx context.Context, req *v1.GetConfigReq) (res *v1.GetConfigRes, err error)
	Run(ctx context.Context, req *v1.RunReq) (res *v1.RunRes, err error)
	Clone(ctx context.Context, req *v1.CloneReq) (res *v1.CloneRes, err error)
}
