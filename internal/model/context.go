package model

import "github.com/casbin/casbin/v2"

type ServiceContext struct {
	CasbinEnforcer *casbin.Enforcer
}

type RequestUser struct {
	UserId   int    `json:"userId"`
	RealName string `json:"realName"`
	Username string `json:"username"`
}
