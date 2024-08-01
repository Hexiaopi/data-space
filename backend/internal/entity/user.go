package entity

import (
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
)

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Avatar     string `json:"avatar"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	State      uint8  `json:"state"`
}

type UserInfo struct {
	User
	Roles       []Role `json:"roles"`
	CurrentRole Role   `json:"current_role"`
}

func ToUser(user model.User) *User {
	return &User{
		ID:         user.ID,
		Name:       user.Name,
		Desc:       user.Desc,
		Avatar:     user.Avatar,
		State:      user.State,
		CreateTime: user.CreateTime.Format(global.DateTimeFormat),
		UpdateTime: user.UpdateTime.Format(global.DateTimeFormat),
	}
}

func ToUsers(users []model.User) []User {
	result := make([]User, 0, len(users))
	for _, user := range users {
		result = append(result, *ToUser(user))
	}
	return result
}
