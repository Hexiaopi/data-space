package model

import "time"

type MenuRole struct {
	MenuId     int64
	RoleId     int64
	CreateTime time.Time
}

func (MenuRole) TableName() string {
	return "sys_menu_role"
}
