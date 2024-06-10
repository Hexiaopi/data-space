package service

import (
	"context"

	"github.com/hexiaopi/data-space/internal/store"
)

type RoleSrv interface {
	Create(ctx context.Context, req *CreateRoleRequest) error
	Update(ctx context.Context, req *UpdateRoleRequest) error
	Delete(ctx context.Context, req *DeleteRoleRequest) error
	List(ctx context.Context, req *ListRoleRequest) (*ListRoleResponse, error)
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

type CreateRoleRequest struct {
}

func (svc *RoleService) Create(ctx context.Context, req *CreateRoleRequest) error {
	return nil
}

type UpdateRoleRequest struct {
}

func (svc *RoleService) Update(ctx context.Context, req *UpdateRoleRequest) error {
	return nil
}

type DeleteRoleRequest struct {
}

func (svc *RoleService) Delete(ctx context.Context, req *DeleteRoleRequest) error {
	return nil
}

type ListRoleRequest struct {
}

type ListRoleResponse struct {
}

func (svc *RoleService) List(ctx context.Context, req *ListRoleRequest) (*ListRoleResponse, error) {
	return nil, nil
}
