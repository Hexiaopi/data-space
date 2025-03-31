package model

import "time"

type Menu struct {
	ID         int64
	Name       string
	Desc       string
	Path       string
	Icon       string
	Component  string
	Type       string
	Order      int64
	ParentId   int64
	CreateTime time.Time
	UpdateTime time.Time
	State      uint8
	Hidden     bool
}

func (Menu) TableName() string {
	return "sys_menu"
}
