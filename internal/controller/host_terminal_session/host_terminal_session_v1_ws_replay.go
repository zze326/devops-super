package host_terminal_session

import (
	"context"
	"devops-super/internal/service"

	"devops-super/api/host_terminal_session/v1"
)

func (c *ControllerV1) WsReplay(ctx context.Context, req *v1.WsReplayReq) (res *v1.WsReplayRes, err error) {
	err = service.HostTerminalSession().Replay(ctx, req.Id)
	return
}
