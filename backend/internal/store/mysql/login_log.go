package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
)

type LoginLogDao struct {
	db *gorm.DB
}

func NewLoginLogDao(db *gorm.DB) *LoginLogDao {
	return &LoginLogDao{db: db}
}

func (dao *LoginLogDao) Create(ctx context.Context, loginLog *model.LoginLog) error {
	if loginLog == nil {
		return nil
	}
	loginLog.CreateTime = time.Now()
	return dao.db.WithContext(ctx).Create(loginLog).Error
}

func (dao *LoginLogDao) List(ctx context.Context, options ...store.Option) ([]model.LoginLog, error) {
	logs := make([]model.LoginLog, 0)
	query := dao.db.WithContext(ctx).Model(&model.LoginLog{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (dao *LoginLogDao) Count(ctx context.Context, options ...store.Option) (int64, error) {
	var count int64
	query := dao.db.WithContext(ctx).Model(&model.LoginLog{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
