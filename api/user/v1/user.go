package v1

import (
	"devops-super/api"
	"devops-super/internal/model/entity"
	"devops-super/internal/model/mid"
	"github.com/gogf/gf/v2/frame/g"
)

type GetPageLstReq struct {
	g.Meta `method:"get" path:"/user/page-list" summary:"分页获取用户列表" tags:"用户"`
	*api.PageLstReq
}

type GetPageLstRes struct {
	*api.PageLstRes[*entity.User]
}

type AddReq struct {
	g.Meta `method:"post" path:"/user" summary:"新增用户" tags:"用户"`
	*mid.User
}

type AddRes struct{}

type UptReq struct {
	g.Meta `method:"put" path:"/user/{id}" summary:"更新用户" tags:"用户"`
	Id     int ` v:"min:1#id必须" path:"id"`
	*mid.User
}

type UptRes struct{}

type UptPasswordReq struct {
	g.Meta   `method:"patch" path:"/user/{id}/password" summary:"更新用户密码" tags:"用户"`
	Id       int    ` v:"min:1#id必须" path:"id"`
	Password string `v:"required|length:6,30#请输入密码|密码长度为:{min}到:{max}位"`
}

type UptPasswordRes struct{}

type UptEnabledReq struct {
	g.Meta  `method:"patch" path:"/user/{id}/enabled" summary:"更新用户密码" tags:"用户"`
	Id      int  `v:"min:1#id必须" path:"id"`
	Enabled bool `v:"required" json:"enabled"`
}

type UptEnabledRes struct{}

type DelReq struct {
	g.Meta `method:"delete" path:"/user/{id}" summary:"删除用户" tags:"用户"`
	Id     int ` v:"min:1#id必须" path:"id"`
}

type DelRes struct{}
