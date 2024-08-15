package service

import (
	"github.com/hexiaopi/data-space/internal/store"
	"github.com/hexiaopi/data-space/pkg/jwt"
	"github.com/hexiaopi/data-space/pkg/logger"
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
	log    logger.Logger
}

func NewService(store store.Factory, option store.Option, jwt jwt.JWT, log logger.Logger) Service {
	return &service{
		store:  store,
		option: option,
		jwt:    jwt,
		log:    log,
	}
}

func (s *service) Users() UserSrv {
	return NewUserService(s.store, s.option, s.jwt, s.log)
}

func (s *service) Roles() RoleSrv {
	return NewRoleService(s.store, s.option, s.log)
}

func (s *service) Menus() MenuSrv {
	return NewMenuService(s.store, s.option, s.log)
}

func (s *service) Departments() DepartmentSrv {
	return NewDepartmentService(s.store, s.option, s.log)
}
