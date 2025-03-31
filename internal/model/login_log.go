package model

import "time"

type LoginLog struct {
	ID           int64
	UserName     string
	RemoteIP     string
	UserAgent    string
	LoginResult  uint8
	ResultDetail string
	CreateTime   time.Time
	UpdateTime   time.Time
}

const (
	LoginResultSuccess = 1
	LoginResultFail    = 2
)

func (LoginLog) TableName() string {
	return "sys_login_log"
}
