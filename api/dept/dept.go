// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dept

import (
	"context"

	"devops-super/api/dept/v1"
)

type IDeptV1 interface {
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	GetLst(ctx context.Context, req *v1.GetLstReq) (res *v1.GetLstRes, err error)
	Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
}
