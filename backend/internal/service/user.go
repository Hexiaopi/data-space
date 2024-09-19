package service

import (
	"context"

	"github.com/hexiaopi/data-space/internal/entity"
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
	"github.com/hexiaopi/data-space/pkg/jwt"
	"github.com/hexiaopi/data-space/pkg/logger"
)

type UserSrv interface {
	Login(ctx context.Context, param *LoginRequest) (*LoginResponse, error)
	Info(ctx context.Context, id int64) (*InfoResponse, error)
	Create(ctx context.Context, param *CreateUserRequest) error
	Update(ctx context.Context, param *UpdateUserRequest) error
	Delete(ctx context.Context, param *DeleteUserRequest) error
	List(ctx context.Context, param *ListUserRequest) (*ListUserResponse, error)
}

type UserService struct {
	store  store.Factory
	option store.Option
	jwt    jwt.JWT
	log    logger.Logger
}

func NewUserService(store store.Factory, option store.Option, jwt jwt.JWT, log logger.Logger) *UserService {
	return &UserService{
		store:  store,
		option: option,
		jwt:    jwt,
		log:    log,
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
		svc.log.Errorf("store user get err: %v", err)
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
		svc.log.Errorf("jwt generate token err: %v", err)
		return nil, global.JWTGenerateTokenFail
	}
	return &LoginResponse{AccessToken: token}, nil
}

type InfoResponse struct {
	entity.UserInfo
}

func (svc *UserService) Info(ctx context.Context, userId int64) (*InfoResponse, error) {
	options := make([]store.Option, 0, 1)
	options = append(options, svc.option.WithId(userId))
	user, err := svc.store.Users().Get(ctx, options...)
	if err != nil {
		svc.log.Errorf("store user get err: %v", err)
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
		svc.log.Errorf("store list user roles err: %v", err)
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
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	Avatar       string `json:"avatar"`
	Password     string `json:"password"`
	State        uint8  `json:"state"`
	DepartmentId int64  `json:"department_id"`
}

func (svc *UserService) Create(ctx context.Context, req *CreateUserRequest) error {
	if err := svc.store.Tx(ctx, func(ctx context.Context, factory store.Factory) error {
		user := &model.User{
			Name:     req.Name,
			Desc:     req.Desc,
			Avatar:   req.Avatar,
			Password: req.Password,
			State:    req.State,
		}
		if err := user.EncryptPassword(); err != nil {
			svc.log.Errorf("user encrypt password err: %v", err)
			return global.UserEncryptPasswordFail
		}
		if err := factory.Users().Create(ctx, user); err != nil {
			svc.log.Errorf("store user create err: %v", err)
			return err
		}
		if req.DepartmentId > 0 {
			if err := factory.Departments().CreateUser(ctx, req.DepartmentId, user.ID); err != nil {
				svc.log.Errorf("store department create user err: %v", err)
				return global.DepartmentCreateUserFail
			}
		}
		return nil
	}); err != nil {
		svc.log.Errorf("store tx err: %v", err)
		return global.UserCreateFail
	}
	return nil
}

type UpdateUserRequest struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	State    uint8  `json:"state"`
}

func (svc *UserService) Update(ctx context.Context, req *UpdateUserRequest) error {
	user := &model.User{
		ID:     req.Id,
		Name:   req.Name,
		Desc:   req.Desc,
		Avatar: req.Avatar,
		State:  req.State,
	}
	if err := user.EncryptPassword(); err != nil {
		svc.log.Errorf("user encrypt password err: %v", err)
		return global.UserEncryptPasswordFail
	}
	if err := svc.store.Users().Update(ctx, user); err != nil {
		svc.log.Errorf("store user update err: %v", err)
		return global.UserUpdateFail
	}
	return nil
}

type DeleteUserRequest struct {
	Id int64
}

func (svc *UserService) Delete(ctx context.Context, req *DeleteUserRequest) error {
	if err := svc.store.Users().Delete(ctx, req.Id); err != nil {
		svc.log.Errorf("store user delete err: %v", err)
		return global.UserDeleteFail
	}
	return nil
}

type ListUserRequest struct {
	Name         string `json:"name"`
	State        uint8  `json:"state"`
	DepartmentId int64  `json:"department_id"`
	PageNum      int    `json:"page_num"`
	PageSize     int    `json:"page_size"`
}

type ListUserResponse struct {
	List  []entity.User `json:"list"`
	Total int64         `json:"total"`
}

func (svc *UserService) List(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	var res ListUserResponse
	options := make([]store.Option, 0)
	if req.Name != "" {
		options = append(options, svc.option.WithName(req.Name))
	}
	if req.State > 0 {
		options = append(options, svc.option.WithState(req.State))
	}
	if req.DepartmentId > 0 {
		options = append(options, svc.option.WithDepartmentId(req.DepartmentId))
	}
	count, err := svc.store.Users().Count(ctx, options...)
	if err != nil {
		svc.log.Errorf("store user count err: %v", err)
		return nil, global.UserCountFail
	}
	if count == 0 {
		return &res, nil
	}
	options = append(options, svc.option.WithPage(req.PageNum, req.PageSize))
	users, err := svc.store.Users().List(ctx, options...)
	if err != nil {
		svc.log.Errorf("store user list err: %v", err)
		return nil, global.UserListFail
	}
	res.List = entity.ToUsers(users)
	res.Total = count
	return &res, nil
}
