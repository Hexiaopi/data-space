package model

import (
	"time"

	"github.com/hexiaopi/data-space/pkg/auth"
)

type User struct {
	ID         int64     `gorm:"column:id"`
	UserName   string    `gorm:"column:username"`
	PassWord   string    `gorm:"column:password"`
	Avatar     string    `gorm:"column:avatar"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	State      uint8     `gorm:"column:state"`
}

func (User) TableName() string {
	return "sys_user"
}

func (u *User) EncryptPassword() error {
	password, err := auth.Encrypt(u.PassWord)
	if err != nil {
		return err
	}
	u.PassWord = password
	return nil
}

func (u *User) ComparePassword(password string) error {
	return auth.Compare(u.PassWord, password)
}
