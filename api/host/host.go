// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package host

import (
	"context"

	"devops-super/api/host/v1"
)

type IHostV1 interface {
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
	GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error)
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	Upt(ctx context.Context, req *v1.UptReq) (res *v1.UptRes, err error)
	Del(ctx context.Context, req *v1.DelReq) (res *v1.DelRes, err error)
	TestSsh(ctx context.Context, req *v1.TestSshReq) (res *v1.TestSshRes, err error)
	DownloadFile(ctx context.Context, req *v1.DownloadFileReq) (res *v1.DownloadFileRes, err error)
	WsTerminal(ctx context.Context, req *v1.WsTerminalReq) (res *v1.WsTerminalRes, err error)
	WsSftpFileManager(ctx context.Context, req *v1.WsSftpFileManagerReq) (res *v1.WsSftpFileManagerRes, err error)
}
