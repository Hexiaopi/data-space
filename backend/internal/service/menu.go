package service

import (
	"context"
	"log"

	"github.com/hexiaopi/data-space/internal/entity"
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
)

type MenuSrv interface {
	AclTree(ctx context.Context, param *AclMenuTreeRequest) (*MenuTreeResponse, error)
	Tree(ctx context.Context, param *MenuTreeRequest) (*MenuTreeResponse, error)
	List(ctx context.Context, param *MenuListRequest) (*MenuListResponse, error)
	Delete(ctx context.Context, id int64) error
	Create(ctx context.Context, param *MenuCreateRequest) error
	Update(ctx context.Context, id int64, param *MenuUpdateRequest) error
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

type AclMenuTreeRequest struct {
	UserId int64
}

type MenuTreeResponse []*entity.Menu

func (svc *MenuService) AclTree(ctx context.Context, param *AclMenuTreeRequest) (*MenuTreeResponse, error) {
	options := make([]store.Option, 0)
	options = append(options, svc.option.WithId(param.UserId))
	user, err := svc.store.Users().Get(ctx, options...)
	if err != nil {
		log.Println(err)
		return nil, global.UserGetFail
	}
	roles, err := svc.store.Roles().ListUserRoles(ctx, user.ID)
	if err != nil {
		log.Println(err)
		return nil, global.RoleListFail
	}
	if len(roles) == 0 {
		return nil, nil
	}
	menus, err := svc.store.Menus().ListRoleMenus(ctx, roles[0].ID)
	if err != nil {
		log.Println(err)
		return nil, global.MenuTreeFail
	}
	tree := entity.GenerateMenuTree(menus)
	res := MenuTreeResponse(tree)
	return &res, nil
}

type MenuTreeRequest struct {
	Name string `json:"name"`
}

func (svc *MenuService) Tree(ctx context.Context, param *MenuTreeRequest) (*MenuTreeResponse, error) {
	options := make([]store.Option, 0)
	if param.Name != "" {
		options = append(options, svc.option.WithName(param.Name))
	}
	options = append(options, svc.option.WithState(model.StateEnable))
	menus, err := svc.store.Menus().List(ctx, options...)
	if err != nil {
		log.Println(err)
		return nil, global.MenuTreeFail
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

type MenuCreateRequest struct {
	entity.Menu
}

func (svc *MenuService) Create(ctx context.Context, param *MenuCreateRequest) error {
	var menu model.Menu
	menu.Name = param.Name
	menu.Desc = param.Desc
	menu.Path = param.Path
	menu.Icon = param.Icon
	menu.Component = param.Component
	menu.Type = param.Type
	menu.Order = param.Order
	menu.ParentId = param.ParentId
	menu.State = model.StateEnable
	if err := svc.store.Menus().Create(ctx, &menu); err != nil {
		log.Println(err)
		return global.MenuCreateFail
	}
	return nil
}

type MenuUpdateRequest struct {
	entity.Menu
}

func (svc *MenuService) Update(ctx context.Context, id int64, param *MenuUpdateRequest) error {
	menu, err := svc.store.Menus().Get(ctx, svc.option.WithId(id))
	if err != nil {
		log.Println(err)
		return global.MenuGetFail
	}
	menu.Name = param.Name
	menu.Desc = param.Desc
	menu.Path = param.Path
	menu.Icon = param.Icon
	menu.Component = param.Component
	menu.Type = param.Type
	menu.Order = param.Order
	if err := svc.store.Menus().Update(ctx, menu); err != nil {
		log.Println(err)
		return global.MenuUpdateFail
	}
	return nil
}

func (svc *MenuService) Delete(ctx context.Context, id int64) error {
	if err := svc.store.Menus().Delete(ctx, id); err != nil {
		log.Println(err)
		return global.MenuDeleteFail
	}
	return nil
}
