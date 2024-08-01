package store

import (
	"context"

	"github.com/hexiaopi/data-space/internal/model"
)

type DepartmentStore interface {
	List(ctx context.Context, options ...Option) ([]model.Department, error)
	Count(ctx context.Context, options ...Option) (int64, error)
	Get(ctx context.Context, options ...Option) (*model.Department, error)
	Create(ctx context.Context, department *model.Department) error
	Update(ctx context.Context, department *model.Department) error
	Delete(ctx context.Context, id int64) error
	CreateUser(ctx context.Context, departmentId, userId int64) error
}
