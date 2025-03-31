package model

import (
	"time"

	"github.com/hexiaopi/data-space/pkg/auth"
)

type User struct {
	ID         int64
	Name       string
	Desc       string
	Password   string
	Avatar     string
	CreateTime time.Time
	UpdateTime time.Time
	State      uint8
}

func (User) TableName() string {
	return "sys_user"
}

func (u *User) EncryptPassword() error {
	password, err := auth.Encrypt(u.Password)
	if err != nil {
		return err
	}
	u.Password = password
	return nil
}

func (u *User) ComparePassword(password string) error {
	return auth.Compare(u.Password, password)
}
