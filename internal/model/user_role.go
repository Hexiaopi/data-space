package model

import "time"

type UserRole struct {
	ID         int64
	UserId     int64
	RoleId     int64
	CreateTime time.Time
}

func (UserRole) TableName() string {
	return "sys_user_role"
}
