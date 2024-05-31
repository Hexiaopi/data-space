package store

import "context"

type TxFunc = func(ctx context.Context, factory Factory) error

type Factory interface {
	Tx(ctx context.Context, fn TxFunc) error
	Close() error
	Users() UserStore
}

var client Factory

func Client() Factory {
	return client
}

func SetClient(f Factory) {
	client = f
}
