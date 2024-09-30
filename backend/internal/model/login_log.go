package model

import "time"

type LoginLog struct {
	ID         int64
	UserId     int64
	RemoteIP   string
	UserAgent  string
	CreateTime time.Time
}

func (LoginLog) TableName() string {
	return "sys_login_log"
}
