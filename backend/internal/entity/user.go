package entity

import (
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
)

type User struct {
	ID         int64  `json:"id"`          //id
	Name       string `json:"name"`        //姓名
	Desc       string `json:"desc"`        //描述
	Avatar     string `json:"avatar"`      //头像
	CreateTime string `json:"create_time"` //创建时间
	UpdateTime string `json:"update_time"` //更新时间
	State      uint8  `json:"state"`       //状态
}

type UserInfo struct {
	User
	Roles       []Role `json:"roles"`        //角色列表
	CurrentRole Role   `json:"current_role"` //当前角色
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
