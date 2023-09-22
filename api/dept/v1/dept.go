package v1

import (
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/frame/g"
)

type AddReq struct {
	g.Meta `path:"/dept" method:"post" tags:"部门" summary:"新增部门"`
	*mid.Dept
}

type AddRes struct{}

type GetLstReq struct {
	g.Meta `path:"/dept/list" method:"get" tags:"部门" summary:"获取部门列表"`
	Search string `p:"search"`
}

type GetLstRes struct {
	List []*entity.Dept `json:"list"`
}

type UptReq struct {
	g.Meta `path:"/dept/{id}" method:"put" tags:"部门" summary:"更新部门"`
	Id     int ` v:"min:1#id必须" path:"id"`
	*mid.Dept
}

type UptRes struct{}

type DelReq struct {
	g.Meta `path:"/dept/{id}" method:"delete" tags:"部门" summary:"删除部门"`
	Id     int `v:"min:1#id必须" path:"id" `
}

type DelRes struct{}
