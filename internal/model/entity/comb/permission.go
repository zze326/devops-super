package comb

import (
	"devops-super/internal/consts"
	"devops-super/internal/model/entity"
)

type Permission struct {
	*entity.Permission
	//Parent   *Permission   `orm:"with:id=parent_id" json:"parent"`
	Children []*Permission `orm:"with:parent_id=id" json:"children"`
}

func (s *Permission) AuthCodes() (codes []string) {
	for _, child := range s.Children {
		if child.Type == consts.PERMISSION_TYPE_ABLE {
			codes = append(codes, child.Name)
		}
	}
	return
}
