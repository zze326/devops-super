package v1

import (
	"devops-super/api"
	"devops-super/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetPageLstReq struct {
	g.Meta `method:"get" path:"/host-terminal-session/page-list" summary:"分页获取主机终端会话记录列表" tags:"主机终端会话"`
	*api.PageLstReq
}

type GetPageLstRes struct {
	*api.PageLstRes[*entity.HostTerminalSession]
}

type CheckSessionFileReq struct {
	g.Meta `method:"get" path:"/host-terminal-session/{id}/check-file" summary:"检查会话文件是否存在" tags:"主机终端会话"`
	Id     int `v:"required" path:"id"`
}

type CheckSessionFileRes struct {
	Exists bool `json:"exists"`
}
