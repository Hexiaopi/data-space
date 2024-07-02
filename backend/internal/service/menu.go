package service

import (
	"context"
	"log"

	"github.com/hexiaopi/data-space/internal/entity"
	"github.com/hexiaopi/data-space/internal/store"
)

type MenuSrv interface {
	Tree(ctx context.Context, param *MenuTreeRequest) (*MenuTreeResponse, error)
	List(ctx context.Context, param *MenuListRequest) (*MenuListResponse, error)
}

type MenuService struct {
	store  store.Factory
	option store.Option
}

func NewMenuService(store store.Factory, option store.Option) *MenuService {
	return &MenuService{
		store:  store,
		option: option,
	}
}

type MenuTreeRequest struct {
	UserId int64
}

type MenuTreeResponse []*entity.Menu

func (svc *MenuService) Tree(ctx context.Context, param *MenuTreeRequest) (*MenuTreeResponse, error) {
	options := make([]store.Option, 0)
	options = append(options, svc.option.WithId(param.UserId))
	user, err := svc.store.Users().Get(ctx, options...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	roles, err := svc.store.Roles().ListUserRoles(ctx, user.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(roles) == 0 {
		return nil, nil
	}
	menus, err := svc.store.Menus().ListRoleMenus(ctx, roles[0].ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	tree := entity.GenerateMenuTree(menus)
	res := MenuTreeResponse(tree)
	return &res, nil
}

type MenuListRequest struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

type MenuListResponse struct {
}

func (svc *MenuService) List(ctx context.Context, param *MenuListRequest) (*MenuListResponse, error) {
	return nil, nil
}
