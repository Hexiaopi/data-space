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

func WithPage(pageSize, pageNum int) Option {
	return func(d *gorm.DB) {
		d.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
}

func WithUserName(username string) Option {
	return func(db *gorm.DB) {
		db.Where("username = ?", username)
	}
}

func WithState(state uint8) Option {
	return func(db *gorm.DB) {
		db.Where("state = ?", state)
	}
}

func (o Option) WithUserName(username string) store.Option {
	return WithUserName(username)
}

func (o Option) WithState(state uint8) store.Option {
	return WithState(state)
}

func (o Option) WithPage(pageSize, pageNum int) store.Option {
	return WithPage(pageSize, pageNum)
}
