package service

import (
	"context"
	log "log/slog"

	"github.com/hexiaopi/data-space/internal/entity"
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/store"
	"github.com/hexiaopi/data-space/pkg/jwt"
)

type UserSrv interface {
	Login(ctx context.Context, param *LoginRequest) (*LoginResponse, error)
	Info(ctx context.Context) (*InfoResponse, error)
	Create(ctx context.Context, param *CreateUserRequest) error
	Update(ctx context.Context, param *UpdateUserRequest) error
	Delete(ctx context.Context, param *DeleteUserRequest) error
	List(ctx context.Context, param *ListUserRequest) (*ListUserResponse, error)
}

type UserService struct {
	store  store.Factory
	option store.Option
	jwt    jwt.JWT
}

func NewUserService(store store.Factory, option store.Option, jwt jwt.JWT) *UserService {
	return &UserService{
		store:  store,
		option: option,
		jwt:    jwt,
	}
}

type LoginRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (svc *UserService) Login(ctx context.Context, param *LoginRequest) (*LoginResponse, error) {
	options := make([]store.Option, 0)
	options = append(options, svc.option.WithName(param.UserName))
	user, err := svc.store.Users().Get(ctx, options...)
	if err != nil {
		log.Error("store user get", err)
		return nil, global.UserGetFail
	}
	if user == nil {
		return nil, global.UserNotExist
	}
	if err := user.ComparePassword(param.PassWord); err != nil {
		return nil, global.UserPassWordWrong
	}
	token, err := svc.jwt.GenerateToken(user.ID, user.Name, user.Password)
	if err != nil {
		log.Error("jwt generate token", err)
		return nil, global.JWTGenerateTokenFail
	}
	return &LoginResponse{AccessToken: token}, nil
}

type InfoResponse struct {
	entity.UserInfo
}

func (svc *UserService) Info(ctx context.Context) (*InfoResponse, error) {
	userId := ctx.Value(global.UserId).(int64)
	options := make([]store.Option, 0, 1)
	options = append(options, svc.option.WithId(userId))
	user, err := svc.store.Users().Get(ctx, options...)
	if err != nil {
		return nil, global.UserGetFail
	}
	if user == nil {
		return nil, global.UserNotExist
	}
	res := InfoResponse{UserInfo: entity.UserInfo{
		User: entity.User{
			ID:         user.ID,
			Name:       user.Name,
			Avatar:     user.Avatar,
			CreateTime: user.CreateTime.Format(global.DateTimeFormat),
			UpdateTime: user.UpdateTime.Format(global.DateTimeFormat),
		},
	}}
	roles, err := svc.store.Roles().ListUserRoles(ctx, userId)
	if err != nil {
		return nil, global.RoleGetFail
	}
	if roles == nil {
		return nil, global.RoleNotExist
	}
	res.Roles = make([]entity.Role, 0, len(roles))
	for _, role := range roles {
		res.Roles = append(res.Roles, entity.Role{
			ID:         role.ID,
			Name:       role.Name,
			CreateTime: role.CreateTime.Format(global.DateTimeFormat),
			UpdateTime: role.UpdateTime.Format(global.DateTimeFormat),
		})
	}
	res.CurrentRole = res.Roles[0]
	return &res, nil
}

type CreateUserRequest struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (svc *UserService) Create(ctx context.Context, param *CreateUserRequest) error {
	return nil
}

type UpdateUserRequest struct {
}

func (svc *UserService) Update(ctx context.Context, param *UpdateUserRequest) error {
	return nil
}

type DeleteUserRequest struct {
	Id int64
}

func (svc *UserService) Delete(ctx context.Context, param *DeleteUserRequest) error {
	return nil
}

type ListUserRequest struct {
}

type ListUserResponse struct {
}

func (svc *UserService) List(ctx context.Context, param *ListUserRequest) (*ListUserResponse, error) {
	return nil, nil
}
