package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type WsLogReq struct {
	g.Meta `method:"get" path:"/ci-pipeline-run/{id}/log" summary:"获取流水线日志" tags:"CI 流水线运行记录"`
	Id     int `v:"required" path:"id"`
}

type WsLogRes struct{}

type WsPageLstReq struct {
	g.Meta `method:"get" path:"/ci-pipeline-run/page-list" summary:"Websocket 分页获取 CI 流水线运行记录列表" tags:"CI 流水线运行记录"`
}

type WsPageLstRes struct{}
