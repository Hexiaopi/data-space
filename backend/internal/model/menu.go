package model

import "time"

type Menu struct {
	ID         int64     `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	Desc       string    `gorm:"column:desc"`
	Path       string    `gorm:"column:path"`
	Icon       string    `gorm:"column:icon"`
	Component  string    `gorm:"column:component"`
	Type       string    `gorm:"column:type"`
	Order      int64     `gorm:"column:order"`
	ParentId   int64     `gorm:"column:parent_id"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	State      uint8     `gorm:"column:state"`
}

func (Menu) TableName() string {
	return "sys_menu"
}
