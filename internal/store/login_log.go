package store

import (
	"context"

	"github.com/hexiaopi/data-space/internal/model"
)

type LoginLogStore interface {
	Create(ctx context.Context, loginLog *model.LoginLog) error
	Count(ctx context.Context, options ...Option) (int64, error)
	List(ctx context.Context, options ...Option) ([]model.LoginLog, error)
}
