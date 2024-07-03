package mysql

import (
	"fmt"

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

func WithName(name string) Option {
	return func(db *gorm.DB) {
		db.Where("name = ?", name)
	}
}

func (o Option) WithName(username string) store.Option {
	return WithName(username)
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

func WithField(field string, value interface{}) Option {
	return func(db *gorm.DB) {
		db.Where(fmt.Sprintf("%s = ?", field), value)
	}
}

func (o Option) WithField(field string, value interface{}) store.Option {
	return WithField(field, value)
}
