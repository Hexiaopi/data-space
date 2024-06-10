package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/hexiaopi/data-space/internal/store"
)

type datastore struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *datastore {
	return &datastore{db: db}
}

func (ds *datastore) Tx(ctx context.Context, f store.TxFunc) error {
	return ds.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		dao := &datastore{db: tx}
		return f(ctx, dao)
	})
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return err
	}
	if db == nil {
		return errors.New("db is nil")
	}
	return db.Close()
}

func (ds *datastore) Users() store.UserStore {
	return NewUserDao(ds.db)
}

func (ds *datastore) Roles() store.RoleStore {
	return NewRoleDao(ds.db)
}

func (ds *datastore) Menus() store.MenuStore {
	return NewMenuDao(ds.db)
}
