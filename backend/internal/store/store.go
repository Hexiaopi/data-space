package store

import (
	"context"
	"errors"
)

//go:generate mockgen -self_package=github.com/hexiaopi/data-space/internal/store -destination mock_store.go -package store github.com/hexiaopi/data-space/internal/store Factory,UserStore,RoleStore,MenuStore

type TxFunc = func(ctx context.Context, factory Factory) error

type Factory interface {
	Tx(ctx context.Context, fn TxFunc) error
	Close() error
	Users() UserStore
	Roles() RoleStore
	Menus() MenuStore
	Departments() DepartmentStore
}

var client Factory

func Client() Factory {
	return client
}

func SetClient(f Factory) {
	client = f
}

var (
	ErrNotFound = errors.New("not found")
)
