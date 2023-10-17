// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ci_pipeline_run

import (
	"context"

	"devops-super/api/ci_pipeline_run/v1"
)

type ICiPipelineRunV1 interface {
	GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error)
	WsLog(ctx context.Context, req *v1.WsLogReq) (res *v1.WsLogRes, err error)
}
