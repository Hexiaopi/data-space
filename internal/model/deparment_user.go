package model

import "time"

type DepartmentUser struct {
	DepartmentId int64
	UserId       int64
	CreateTime   time.Time
}

func (DepartmentUser) TableName() string {
	return "sys_department_user"
}
