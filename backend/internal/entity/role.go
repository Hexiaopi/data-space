package entity

import (
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
)

type Role struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	State      uint8  `json:"state"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func ToRole(role model.Role) Role {
	return Role{
		ID:         role.ID,
		Name:       role.Name,
		Desc:       role.Desc,
		State:      role.State,
		CreateTime: role.CreateTime.Format(global.DateTimeFormat),
		UpdateTime: role.UpdateTime.Format(global.DateTimeFormat),
	}
}

func ToRoles(roles []model.Role) []Role {
	result := make([]Role, 0, len(roles))
	for _, role := range roles {
		result = append(result, ToRole(role))
	}
	return result
}

type RoleStatic struct {
	Role
	UserCount int64 `json:"user_count"` // 用户数量
}
