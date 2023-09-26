package host

import (
	"context"
	"devops-super/internal/model/mid"
	"devops-super/internal/service"

	"devops-super/api/host/v1"
)

func (c *ControllerV1) DownloadFile(ctx context.Context, req *v1.DownloadFileReq) (res *v1.DownloadFileRes, err error) {
	err = service.Host().DownloadFile(ctx, &mid.DownloadFileIn{
		Id:   req.Id,
		Path: req.Path,
	})
	return
}
