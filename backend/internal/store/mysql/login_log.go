package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/data-space/internal/model"
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
