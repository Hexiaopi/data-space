package entity

import (
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
)

type Role struct {
	ID         int64  `json:"id"`          //id
	Name       string `json:"name"`        //名称
	Desc       string `json:"desc"`        //描述
	State      uint8  `json:"state"`       //状态
	CreateTime string `json:"create_time"` //创建时间
	UpdateTime string `json:"update_time"` //更新时间
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
