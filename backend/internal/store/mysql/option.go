package mysql

import (
	"gorm.io/gorm"

	"github.com/hexiaopi/data-space/internal/store"
)

func NewOption() store.Option {
	var option store.Option = Option(nil)
	return option
}

type Option func(*gorm.DB)

func WithPage(pageNum, pageSize int) Option {
	return func(d *gorm.DB) {
		d.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
}

func (o Option) WithPage(pageNum, pageSize int) store.Option {
	return WithPage(pageNum, pageSize)
}

func WithUserName(username string) Option {
	return func(db *gorm.DB) {
		db.Where("name = ?", username)
	}
}

func (o Option) WithUserName(username string) store.Option {
	return WithUserName(username)
}

func WithId(id int64) Option {
	return func(db *gorm.DB) {
		db.Where("id = ?", id)
	}
}

func (o Option) WithId(id int64) store.Option {
	return WithId(id)
}

func WithState(state uint8) Option {
	return func(db *gorm.DB) {
		db.Where("state = ?", state)
	}
}

func (o Option) WithState(state uint8) store.Option {
	return WithState(state)
}
