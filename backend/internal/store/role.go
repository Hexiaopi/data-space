package store

import (
	"context"

	"github.com/hexiaopi/data-space/internal/model"
)

type RoleStore interface {
	Get(ctx context.Context, options ...Option) (*model.Role, error)
	Create(ctx context.Context, role *model.Role) error
	Update(ctx context.Context, role *model.Role) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, options ...Option) ([]model.Role, error)
	Count(ctx context.Context, options ...Option) (int64, error)
	ListUserRoles(ctx context.Context, userId int64, options ...Option) ([]model.Role, error)
}
