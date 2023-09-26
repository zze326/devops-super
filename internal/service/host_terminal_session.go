// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"devops-super/api"
	"devops-super/internal/model/do"
	"devops-super/internal/model/entity"
)

type (
	IHostTerminalSession interface {
		Get(ctx context.Context, in *do.HostTerminalSession) (out *entity.HostTerminalSession, err error)
		GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.HostTerminalSession], err error)
		CheckSessionFile(ctx context.Context, id int) (bool, error)
		Replay(ctx context.Context, id int) error
	}
)

var (
	localHostTerminalSession IHostTerminalSession
)

func HostTerminalSession() IHostTerminalSession {
	if localHostTerminalSession == nil {
		panic("implement not found for interface IHostTerminalSession, forgot register?")
	}
	return localHostTerminalSession
}

func RegisterHostTerminalSession(i IHostTerminalSession) {
	localHostTerminalSession = i
}
