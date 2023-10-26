package v1

import "github.com/gogf/gf/v2/frame/g"

type CancelReq struct {
	g.Meta `method:"delete" path:"/ci-pipeline-run/{id}/cancel" summary:"取消流水线的执行" tags:"CI 流水线运行记录"`
	Id     int ` v:"min:1#id必须" path:"id"`
}

type CancelRes struct{}
