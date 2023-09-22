package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" tags:"用户" method:"post" summary:"登录"`
	Username string `v:"required#请输入用户名" json:"username" dc:"用户名"`
	Password string `v:"required#请输入密码" json:"password" dc:"密码"`
}

type LoginRes struct {
	Username     string   `json:"username"`
	RealName     string   `json:"realName"`
	Token        string   `json:"token"`
	Expires      int64    `json:"expires"`
	RefreshAfter int64    `json:"refreshAfter"`
	Roles        []string `json:"roles"`
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh-token" method:"post" tags:"用户" summary:"刷新 Token"`
}

type RefreshTokenRes struct {
	Username     string   `json:"username"`
	RealName     string   `json:"realName"`
	Token        string   `json:"token"`
	Expire       int64    `json:"expire"`
	RefreshAfter int64    `json:"refreshAfter"`
	Roles        []string `json:"roles"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"用户" summary:"用户登出"`
}

type LogoutRes struct{}
