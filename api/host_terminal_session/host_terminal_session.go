// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package host_terminal_session

import (
	"context"

	"devops-super/api/host_terminal_session/v1"
)

type IHostTerminalSessionV1 interface {
	GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error)
	CheckSessionFile(ctx context.Context, req *v1.CheckSessionFileReq) (res *v1.CheckSessionFileRes, err error)
	WsReplay(ctx context.Context, req *v1.WsReplayReq) (res *v1.WsReplayRes, err error)
}
