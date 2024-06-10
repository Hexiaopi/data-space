package store

import (
	"context"

	"github.com/hexiaopi/data-space/internal/model"
)

type MenuStore interface {
	// Get(ctx context.Context, options ...Option) (*model.Menu, error)
	// Create(ctx context.Context, menu *model.Menu) error
	// Update(ctx context.Context, menu *model.Menu) error
	// Delete(ctx context.Context, id int64) error
	// List(ctx context.Context, options ...Option) ([]model.Menu, error)
	// Count(ctx context.Context, options ...Option) (int64, error)
	ListRoleMenus(ctx context.Context, roleId int64, options ...Option) ([]model.Menu, error)
}
