// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package host_group

import (
	"context"

	"devops-super/api/host_group/v1"
)

type IHostGroupV1 interface {
	GetLst(ctx context.Context, req *v1.GetLstReq) (res *v1.GetLstRes, err error)
	GetPartialList(ctx context.Context, req *v1.GetPartialListReq) (res *v1.GetPartialListRes, err error)
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
}
