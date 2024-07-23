package model

import "time"

type Department struct {
	ID         int64
	Name       string
	Desc       string
	CreateTime time.Time
	UpdateTime time.Time
	State      uint8
}

func (Department) TableName() string {
	return "sys_department"
}
