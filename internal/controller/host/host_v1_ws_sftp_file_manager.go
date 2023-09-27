package host

import (
	"context"
	"devops-super/internal/model/do"
	"devops-super/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"

	"devops-super/api/host/v1"
)

func (c *ControllerV1) WsSftpFileManager(ctx context.Context, req *v1.WsSftpFileManagerReq) (res *v1.WsSftpFileManagerRes, err error) {
	eHost, err := service.Host().Get(ctx, &do.Host{Id: req.Id})
	if err != nil {
		return nil, err
	}
	if eHost == nil {
		return nil, gerror.New("主机不存在")
	}

	can, err := service.Host().CanAccess(ctx, eHost)
	if err != nil {
		return nil, err
	}

	if !can {
		return nil, gerror.New("未授权")
	}

	err = service.Host().WsSftpFileManager(ctx, eHost)
	return
}
