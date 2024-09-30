package store

import (
	"context"

	"github.com/hexiaopi/data-space/internal/model"
)

type LoginLogStore interface {
	Create(ctx context.Context, loginLog *model.LoginLog) error
}
