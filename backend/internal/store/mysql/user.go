package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) Get(ctx context.Context, options ...store.Option) (*model.User, error) {
	user := model.User{}
	query := dao.db.WithContext(ctx).Model(&model.User{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDao) Create(ctx context.Context, user *model.User) error {
	if user == nil {
		return nil
	}
	now := time.Now()
	user.CreateTime = now
	user.UpdateTime = now
	return dao.db.WithContext(ctx).Create(user).Error
}

func (dao *UserDao) Update(ctx context.Context, user *model.User) error {
	return dao.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"name":        user.Name,
			"avatar":      user.Avatar,
			"state":       user.State,
			"update_time": time.Now(),
		}).Error
}

func (dao *UserDao) Delete(ctx context.Context, id int64) error {
	return dao.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"state":       model.StateDelete,
			"update_time": time.Now(),
		}).Error
}

func (dao *UserDao) List(ctx context.Context, options ...store.Option) ([]model.User, error) {
	users := make([]model.User, 0)
	query := dao.db.WithContext(ctx).Model(&model.User{}).Joins("left join sys_department_user on sys_user.id = sys_department_user.user_id")
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *UserDao) Count(ctx context.Context, options ...store.Option) (int64, error) {
	var count int64
	query := dao.db.WithContext(ctx).Model(&model.User{}).Joins("left join sys_department_user on sys_user.id = sys_department_user.user_id")
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
