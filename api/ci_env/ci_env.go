// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ci_env

import (
	"context"

	"devops-super/api/ci_env/v1"
)

type ICiEnvV1 interface {
	GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error)
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
	GetLst(ctx context.Context, req *v1.GetLstReq) (res *v1.GetLstRes, err error)
}
