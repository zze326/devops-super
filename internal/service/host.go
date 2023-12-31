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
	"devops-super/internal/model/mid"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type (
	IHost interface {
		Add(ctx context.Context, in *entity.Host) (err error)
		Upt(ctx context.Context, in *do.Host) (err error)
		GetCountByHostGroupId(ctx context.Context, hostGroupId int) (int, error)
		GetPageLst(ctx context.Context, in *api.PageLstReq) (out *api.PageLstRes[*entity.Host], err error)
		Get(ctx context.Context, in *do.Host) (out *entity.Host, err error)
		Del(ctx context.Context, in *do.Host) (err error)
		TestSSH(ctx context.Context, in *entity.Host) (err error)
		DownloadFile(ctx context.Context, in *mid.DownloadFileIn) error
		CanAccess(ctx context.Context, in *entity.Host) (bool, error)
		GetAuthorizedLst(ctx context.Context) (out []*entity.Host, err error)
		WsSftpFileManager(ctx context.Context, in *entity.Host) (err error)
		SftpClient(in *entity.Host) (*sftp.Client, error)
		WsTerminal(ctx context.Context, in *entity.Host) error
		SshClient(in *entity.Host) (*ssh.Client, error)
	}
)

var (
	localHost IHost
)

func Host() IHost {
	if localHost == nil {
		panic("implement not found for interface IHost, forgot register?")
	}
	return localHost
}

func RegisterHost(i IHost) {
	localHost = i
}
