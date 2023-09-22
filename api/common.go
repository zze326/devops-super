package api

import (
	"github.com/gogf/gf/v2/encoding/gjson"
)

type PageLstReq struct {
	Page     int         `p:"page" v:"page @integer|min:1#页码必填" dc:"页码"`              // 页码
	PageSize int         `p:"pageSize" v:"pageSize @integer|min:1#每页数量必填"  dc:"每页数量"` // 每页数量
	Search   string      `p:"search" dc:"模糊搜索内容"`                                     // 搜索内容
	Wheres   *gjson.Json `p:"wheres" dc:"搜索条件"`
}

func (r *PageLstReq) Offset() int {
	return (r.Page - 1) * r.PageSize
}

func (r *PageLstReq) Limit() int {
	return r.PageSize
}

func (r *PageLstReq) SearchStr() string {
	return "%" + r.Search + "%"
}

type PageLstRes[T any] struct {
	Total int `json:"total" dc:"总数"` // 总数
	List  []T `json:"list" dc:"列表"`  // 列表
}
