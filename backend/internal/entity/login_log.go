package entity

import (
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
)

type LoginLog struct {
	Id           int64  `json:"id"`            //id
	UserName     string `json:"user_name"`     //用户名
	RemoteIp     string `json:"remote_ip"`     //访问ip
	UserAgent    string `json:"user_agent"`    //用户代理
	LoginResult  uint8  `json:"login_result"`  //登录结果
	ResultDetail string `json:"result_detail"` //结果详情
	CreateTime   string `json:"create_time"`   //创建时间
	UpdateTime   string `json:"update_time"`   //更新时间
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
