package v1

import (
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/frame/g"
)

type GetLstReq struct {
	g.Meta `method:"get" path:"/host-group/list" summary:"获取主机组列表" tags:"主机组"`
	Search string `p:"search"`
}

type GetLstRes struct {
	List []*entity.HostGroup `json:"list"`
}

type GetPartialListReq struct {
	g.Meta `method:"get" path:"/host-group/partial-list" summary:"获取主机组列表(部分字段)" tags:"主机组"`
}

type GetPartialListRes struct {
	List []*mid.HostGroupPartial `json:"list"`
}

type AddReq struct {
	g.Meta `method:"post" path:"/host-group" summary:"添加主机组" tags:"主机组"`
	*mid.HostGroup
}

type AddRes struct{}

type UptReq struct {
	g.Meta `method:"put" path:"/host-group/{id}" summary:"更新主机组" tags:"主机组"`
	Id     int ` v:"min:1#id必须" path:"id"`
	*mid.HostGroup
}

type UptRes struct{}

type DelReq struct {
	g.Meta `method:"delete" path:"/host-group/{id}" summary:"删除主机组" tags:"主机组"`
	Id     int ` v:"min:1#id必须" path:"id"`
}

type DelRes struct{}
