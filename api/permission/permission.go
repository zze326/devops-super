// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package permission

import (
	"context"
	
	"devops-super/api/permission/v1"
)

type IPermissionV1 interface {
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	GetLst(ctx context.Context, req *v1.GetLstReq) (res *v1.GetLstRes, err error)
	Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error)
	UptShowLink(ctx context.Context, req *v1.UptShowLinkReq) (res *v1.UptShowLinkRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
	GetRouteLst(ctx context.Context, req *v1.GetRouteLstReq) (res *v1.GetRouteLstRes, err error)
}


