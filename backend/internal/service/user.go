package service

import (
	"context"
	log "log/slog"

	"github.com/hexiaopi/data-space/internal/pkg/retcode"
	"github.com/hexiaopi/data-space/internal/store"
	"github.com/hexiaopi/data-space/pkg/jwt"
)

type UserSrv interface {
	Login(ctx context.Context, param *LoginRequest) (*LoginResponse, error)
	// List(ctx context.Context, param *ListUserRequest) ([]*entity.User, error)
	// Create(ctx context.Context, param *CreateUserRequest) error
	// Update(ctx context.Context, param *UpdateUserRequest) error
	// Delete(ctx context.Context, param *DeleteUserRequest) error
}

type LoginRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type ListUserRequest struct {
}

type CreateUserRequest struct {
}

type UpdateUserRequest struct {
}

type DeleteUserRequest struct {
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

func (svc *UserService) Login(ctx context.Context, param *LoginRequest) (*LoginResponse, error) {
	options := make([]store.Option, 0)
	options = append(options, svc.option.WithUserName(param.UserName))
	user, err := svc.store.Users().Get(ctx, options...)
	if err != nil {
		log.Error("store user get", err)
		return nil, retcode.UserGetFail
	}
	if user == nil {
		return nil, retcode.UserNotExist
	}
	if err := user.ComparePassword(param.PassWord); err != nil {
		return nil, retcode.UserPassWordWrong
	}
	token, err := svc.jwt.GenerateToken(user.ID, user.UserName, user.PassWord)
	if err != nil {
		log.Error("jwt generate token", err)
		return nil, retcode.JWTGenerateTokenFail
	}
	return &LoginResponse{AccessToken: token}, nil
}
