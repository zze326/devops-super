package ci_pipeline_run

import (
	"context"
	"devops-super/api"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gorilla/websocket"
)

func (s *sCiPipelineRun) WsPageLst(ctx context.Context) error {
	ws, err := g.RequestFromCtx(ctx).WebSocket()
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		in := new(api.PageLstReq)
		if err = ws.ReadJSON(in); err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				// 连接异常关闭
				glog.Errorf(ctx, "连接异常关闭: %v", err)
				return err
			}
			break
		}

		out, err := s.GetPageLst(ctx, in)
		if err != nil {
			return err
		}
		if err = ws.WriteJSON(out); err != nil {
			return err
		}
	}

	glog.Info(ctx, "分页连接正常断开")
	return nil
}
