package v1

import "github.com/gogf/gf/v2/frame/g"

type PingReq struct {
	g.Meta `path:"/ping" method:"get" tags:"健康检查" summary:"ping"`
}

type PingRes struct{}
