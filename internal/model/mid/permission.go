package mid

import (
	"github.com/gogf/gf/v2/encoding/gjson"
)

type Permission struct {
	Title      string      `v:"required" json:"title"`
	Name       string      `v:"required" json:"name"`
	Type       int         `v:"required" json:"type"`
	FRoute     string      `json:"fRoute"`
	BRoutes    *gjson.Json `json:"bRoutes"`
	Redirect   string      `json:"redirect"`
	Icon       string      `json:"icon"`
	Rank       int         `json:"rank"`
	ShowLink   bool        `v:"required" json:"showLink"`
	ShowParent bool        `v:"required" json:"showParent"`
	KeepAlive  bool        `v:"required" json:"keepAlive"`
	ParentId   int         `json:"parentId"`
}

type Route struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Title      string   `json:"title"`
	FRoute     string   `json:"fRoute"`
	Redirect   string   `json:"redirect"`
	Icon       string   `json:"icon,omitempty"`
	ParentId   int      `json:"parentId"`
	Rank       int      `json:"rank"`
	ShowLink   bool     `json:"showLink"`
	ShowParent bool     `json:"showParent"`
	Auths      []string `json:"auths,omitempty"`
	KeepAlive  bool     `json:"keepAlive"`
}

type FrontendRouteList []*Route

func (s FrontendRouteList) Len() int { return len(s) }

func (s FrontendRouteList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s FrontendRouteList) Less(i, j int) bool { return s[i].Rank < s[j].Rank }
