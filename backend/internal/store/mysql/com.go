package mysql

import (
	"context"

	"gorm.io/gorm"

	"github.com/hexiaopi/data-space/internal/store"
)

type ComDao struct {
	db *gorm.DB
}

func NewComDao(db *gorm.DB) *ComDao {
	return &ComDao{
		db: db,
	}
}

func (dao *ComDao) Get(ctx context.Context, table string, in interface{}, options ...Option) error {
	query := dao.db.WithContext(ctx).Table(table)
	for _, option := range options {
		option(query)
	}
	err := query.First(in).Error
	if err == gorm.ErrRecordNotFound {
		return store.ErrNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

func (dao *ComDao) Create(ctx context.Context, in interface{}) error {
	return dao.db.WithContext(ctx).Create(in).Error
}

func (dao *ComDao) Update(ctx context.Context, in interface{}) error {
	return dao.db.WithContext(ctx).Save(in).Error
}

func (dao *ComDao) Delete(ctx context.Context, in interface{}) error {
	return dao.db.WithContext(ctx).Delete(in).Error
}

func (dao *ComDao) List(ctx context.Context, table string, in interface{}, options ...Option) error {
	query := dao.db.WithContext(ctx).Table(table)
	for _, option := range options {
		option(query)
	}
	if err := query.Find(in).Error; err != nil {
		return err
	}
	return nil
}
