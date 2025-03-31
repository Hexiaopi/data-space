package service

import (
	"context"

	"github.com/hexiaopi/data-space/internal/entity"
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
	"github.com/hexiaopi/data-space/pkg/logger"
)

type LoginLogSrv interface {
	Create(ctx context.Context, loginLog *model.LoginLog) error
	List(ctx context.Context, req *LoginLogListRequest) (*LoginLogListResponse, error)
}

type LoginLogService struct {
	store  store.Factory
	option store.Option
	log    logger.Logger
}

func NewLoginLogService(store store.Factory, option store.Option, log logger.Logger) *LoginLogService {
	return &LoginLogService{
		store:  store,
		option: option,
		log:    log,
	}
}

func (srv *LoginLogService) Create(ctx context.Context, loginLog *model.LoginLog) error {
	if err := srv.store.LoginLogs().Create(ctx, loginLog); err != nil {
		return err
	}
	return nil
}

type LoginLogListRequest struct {
	UserName    string `json:"user_name"`    //用户名称
	LoginResult int    `json:"login_result"` //登录结果
	PageNum     int    `json:"page_num"`     //页码
	PageSize    int    `json:"page_size"`    //页大小
}

type LoginLogListResponse struct {
	List  []entity.LoginLog `json:"list"`
	Total int64             `json:"total"`
}

func (svc *LoginLogService) List(ctx context.Context, req *LoginLogListRequest) (*LoginLogListResponse, error) {
	var res LoginLogListResponse
	options := make([]store.Option, 0)
	if req.UserName != "" {
		options = append(options, svc.option.WithLikeField("user_name", req.UserName))
	}
	if req.LoginResult > 0 {
		options = append(options, svc.option.WithField("login_result", req.LoginResult))
	}
	count, err := svc.store.LoginLogs().Count(ctx, options...)
	if err != nil {
		svc.log.Errorf("store login log count err: %v", err)
		return nil, global.UserCountFail
	}
	if count == 0 {
		return &res, nil
	}
	options = append(options, svc.option.WithPage(req.PageNum, req.PageSize))
	users, err := svc.store.LoginLogs().List(ctx, options...)
	if err != nil {
		svc.log.Errorf("store login log list err: %v", err)
		return nil, global.UserListFail
	}
	res.List = entity.ToLoginLogs(users)
	res.Total = count
	return &res, nil
}
