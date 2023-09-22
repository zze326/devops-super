package mid

import "github.com/gogf/gf/v2/os/gtime"

type Dept struct {
	Name      string      `v:"required" json:"name"`
	Rank      int         `v:"required" json:"rank"`
	ParentId  int         `json:"parentId"`
	CreatedAt *gtime.Time `json:"createdAt"`
}
