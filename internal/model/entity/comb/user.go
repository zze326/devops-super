package comb

import "devops-super/internal/model/entity"

type User struct {
	*entity.User
	Roles []*entity.Role
}

func (u *User) RoleCodes() []string {
	codes := make([]string, 0)
	for _, role := range u.Roles {
		codes = append(codes, role.Code)
	}
	return codes
}
