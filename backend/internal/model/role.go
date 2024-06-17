package model

import "time"

type Role struct {
	ID         int64
	Name       string
	Desc       string
	CreateTime time.Time
	UpdateTime time.Time
	State      uint8
}

func (Role) TableName() string {
	return "sys_role"
}
