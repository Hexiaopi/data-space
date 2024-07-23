package service

import (
	"github.com/hexiaopi/data-space/internal/store"
	"github.com/hexiaopi/data-space/pkg/jwt"
)

type Service interface {
	Users() UserSrv
	Roles() RoleSrv
	Menus() MenuSrv
	Departments() DepartmentSrv
}

type service struct {
	store  store.Factory
	option store.Option
	jwt    jwt.JWT
}

func NewService(store store.Factory, option store.Option, jwt jwt.JWT) Service {
	return &service{
		store:  store,
		option: option,
		jwt:    jwt,
	}
}

func (s *service) Users() UserSrv {
	return NewUserService(s.store, s.option, s.jwt)
}

func (s *service) Roles() RoleSrv {
	return NewRoleService(s.store, s.option)
}

func (s *service) Menus() MenuSrv {
	return NewMenuService(s.store, s.option)
}

func (s *service) Departments() DepartmentSrv {
	return NewDepartmentService(s.store, s.option)
}
