package comb

import "devops-super/internal/model/entity"

type Role struct {
	*entity.Role
	Permissions []*Permission
}
