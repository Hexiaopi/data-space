package store

import (
	"context"

	"github.com/hexiaopi/data-space/internal/model"
)

type UserStore interface {
	Get(ctx context.Context, options ...Option) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
	Count(ctx context.Context, options ...Option) (int64, error)
	List(ctx context.Context, options ...Option) ([]model.User, error)
}
