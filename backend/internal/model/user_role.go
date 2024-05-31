package model

import "time"

type UserRole struct {
	ID        int64     `gorm:"column:id"`
	UserId    int64     `gorm:"column:user_id"`
	RoleId    int64     `gorm:"column:role_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (UserRole) TableName() string {
	return "sys_user_role"
}
