package host_terminal_session

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/host_terminal_session/v1"
)

func (c *ControllerV1) CheckSessionFile(ctx context.Context, req *v1.CheckSessionFileReq) (res *v1.CheckSessionFileRes, err error) {
	res = new(v1.CheckSessionFileRes)
	res.Exists, err = service.HostTerminalSession().CheckSessionFile(ctx, req.Id)
	return
}
