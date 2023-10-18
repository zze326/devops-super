// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ci_pipeline_run

import (
	"context"

	"devops-super/api/ci_pipeline_run/v1"
)

type ICiPipelineRunV1 interface {
	WsLog(ctx context.Context, req *v1.WsLogReq) (res *v1.WsLogRes, err error)
	WsPageLst(ctx context.Context, req *v1.WsPageLstReq) (res *v1.WsPageLstRes, err error)
}
