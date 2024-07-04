package service

import (
	"context"
	"log"

	"github.com/hexiaopi/data-space/internal/entity"
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
)

type RoleSrv interface {
	Create(ctx context.Context, req *RoleCreateRequest) error
	Update(ctx context.Context, id int64, req *RoleUpdateRequest) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, req *RoleListRequest) (*RoleListResponse, error)
	Get(ctx context.Context, id int64) (*RoleGetResponse, error)
	GetMenu(ctx context.Context, id int64) (*entity.MenuTree, error)
	UpdateMenu(ctx context.Context, id int64, param *UpdateMenuRequest) error
}

type RoleService struct {
	store  store.Factory
	option store.Option
}

func NewRoleService(store store.Factory, option store.Option) *RoleService {
	return &RoleService{
		store:  store,
		option: option,
	}
}

type RoleCreateRequest struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	State uint8  `json:"state"`
}

func (svc *RoleService) Create(ctx context.Context, req *RoleCreateRequest) error {
	var role model.Role
	role.Name = req.Name
	role.Desc = req.Desc
	role.State = req.State
	if err := svc.store.Roles().Create(ctx, &role); err != nil {
		log.Println(err)
		return global.RoleCreateFail
	}
	return nil
}

type RoleUpdateRequest struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	State uint8  `json:"state"`
}

func (svc *RoleService) Update(ctx context.Context, id int64, req *RoleUpdateRequest) error {
	role, err := svc.store.Roles().Get(ctx, svc.option.WithId(id))
	if err != nil {
		log.Println(err)
		return global.RoleGetFail
	}
	role.Name = req.Name
	role.Desc = req.Desc
	role.State = req.State
	if err := svc.store.Roles().Update(ctx, role); err != nil {
		log.Println(err)
		return global.RoleUpdateFail
	}
	return nil
}

func (svc *RoleService) Delete(ctx context.Context, id int64) error {
	if err := svc.store.Roles().Delete(ctx, id); err != nil {
		log.Println(err)
		return global.RoleDeleteFail
	}
	return nil
}

type RoleListRequest struct {
	Name     string `json:"name"`
	State    uint8  `json:"state"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
}

type RoleListResponse struct {
	Roles []entity.Role `json:"roles"`
	Total int64         `json:"total"`
}

func (svc *RoleService) List(ctx context.Context, req *RoleListRequest) (*RoleListResponse, error) {
	var res RoleListResponse
	options := make([]store.Option, 0)
	if req.Name != "" {
		options = append(options, svc.option.WithName(req.Name))
	}
	if req.State > 0 {
		options = append(options, svc.option.WithState(req.State))
	}
	count, err := svc.store.Roles().Count(ctx, options...)
	if err != nil {
		log.Println(err)
		return nil, global.RoleCountFail
	}
	if count == 0 {
		return &res, nil
	}
	options = append(options, svc.option.WithPage(req.PageNum, req.PageSize))
	roles, err := svc.store.Roles().List(ctx, options...)
	if err != nil {
		log.Println(err)
		return nil, global.RoleListFail
	}
	res.Roles = entity.ToRoles(roles)
	res.Total = count
	return &res, nil
}

type RoleGetResponse struct {
	Role  entity.Role     `json:"role"`
	Menus entity.MenuTree `json:"menus"`
}

func (svc *RoleService) Get(ctx context.Context, id int64) (*RoleGetResponse, error) {
	var res RoleGetResponse
	options := make([]store.Option, 0)
	options = append(options, svc.option.WithId(id))
	role, err := svc.store.Roles().Get(ctx, options...)
	if err != nil {
		log.Println(err)
		return nil, global.RoleGetFail
	}
	res.Role = entity.ToRole(*role)

	menus, err := svc.store.Menus().ListRoleMenus(ctx, id, svc.option.WithState(model.StateEnable))
	if err != nil {
		log.Println(err)
		return nil, global.MenuTreeFail
	}

	tree := entity.GenerateMenuTree(menus)
	res.Menus = tree
	return &res, nil
}

type UpdateMenuRequest struct {
	MenuIds []int64 `json:"menu_ids"`
}

func (svc *RoleService) UpdateMenu(ctx context.Context, id int64, param *UpdateMenuRequest) error {
	return svc.store.Tx(ctx, func(ctx context.Context, factory store.Factory) error {
		if err := factory.Roles().DeleteMenus(ctx, id); err != nil {
			log.Println(err)
			return global.RoleDeleteMenuFail
		}
		if len(param.MenuIds) > 0 {
			if err := factory.Roles().CreateMenus(ctx, id, param.MenuIds...); err != nil {
				log.Println(err)
				return global.RoleCreateMenuFail
			}
		}
		return nil
	})
}

func (svc *RoleService) GetMenu(ctx context.Context, id int64) (*entity.MenuTree, error) {
	menus, err := svc.store.Menus().ListRoleMenus(ctx, id, svc.option.WithState(model.StateEnable))
	if err != nil {
		log.Println(err)
		return nil, global.MenuTreeFail
	}
	tree := entity.GenerateMenuTree(menus)
	return &tree, nil
}
