package host_terminal_session

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/host_terminal_session/v1"
)

func (c *ControllerV1) GetPageLst(ctx context.Context, req *v1.GetPageLstReq) (res *v1.GetPageLstRes, err error) {
	res = new(v1.GetPageLstRes)
	res.PageLstRes, err = service.HostTerminalSession().GetPageLst(ctx, req.PageLstReq)
	return
}
