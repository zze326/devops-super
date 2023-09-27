package mid

import "github.com/gogf/gf/v2/encoding/gjson"

type HostGroup struct {
	Name     string      `v:"required" json:"name"`
	Rank     int         `v:"required" json:"rank"`
	ParentId int         `json:"parentId"`
	RoleIds  *gjson.Json `json:"roleIds"`
	UserIds  *gjson.Json `json:"userIds"`
}

type HostGroupPartial struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ParentId  int    `json:"parentId"`
	HostCount int    `json:"hostCount"`
}
