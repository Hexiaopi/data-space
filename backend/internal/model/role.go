package model

import "time"

type Role struct {
	ID         int64     `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	State      uint8     `gorm:"column:state"`
}

func (Role) TableName() string {
	return "sys_role"
}
