package v1

import "github.com/gogf/gf/v2/frame/g"

type WsReplayReq struct {
	g.Meta `method:"get" path:"/host-terminal-session/{id}/replay" summary:"连接终端" tags:"主机"`
	Id     int `v:"required" path:"id"`
}

type WsReplayRes struct{}
