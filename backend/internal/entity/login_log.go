package entity

import (
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
)

type LoginLog struct {
	Id           int64  `json:"id"`
	UserName     string `json:"user_name"`
	RemoteIp     string `json:"remote_ip"`
	UserAgent    string `json:"user_agent"`
	LoginResult  uint8  `json:"login_result"`
	ResultDetail string `json:"result_detail"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

func ToLoginLog(loginLog model.LoginLog) LoginLog {
	return LoginLog{
		Id:           loginLog.ID,
		UserName:     loginLog.UserName,
		RemoteIp:     loginLog.RemoteIP,
		UserAgent:    loginLog.UserAgent,
		LoginResult:  loginLog.LoginResult,
		ResultDetail: loginLog.ResultDetail,
		CreateTime:   loginLog.CreateTime.Format(global.DateTimeFormat),
		UpdateTime:   loginLog.UpdateTime.Format(global.DateTimeFormat),
	}
}

func ToLoginLogs(loginLogs []model.LoginLog) []LoginLog {
	res := make([]LoginLog, 0)
	for _, loginLog := range loginLogs {
		res = append(res, ToLoginLog(loginLog))
	}
	return res
}
