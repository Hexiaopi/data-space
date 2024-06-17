package mysql

import (
	"context"
	"time"

	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{db: db}
}

func (dao *RoleDao) Get(ctx context.Context, options ...store.Option) (*model.Role, error) {
	role := model.Role{}
	query := dao.db.WithContext(ctx).Model(&model.Role{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (dao *RoleDao) Create(ctx context.Context, role *model.Role) error {
	if role == nil {
		return nil
	}
	now := time.Now()
	role.CreateTime = now
	role.UpdateTime = now
	return dao.db.WithContext(ctx).Create(role).Error
}

func (dao *RoleDao) Update(ctx context.Context, role *model.Role) error {
	return dao.db.WithContext(ctx).Model(&model.Role{}).
		Where("id = ?", role.ID).
		Updates(map[string]interface{}{
			"name":        role.Name,
			"desc":        role.Desc,
			"state":       role.State,
			"update_time": time.Now(),
		}).Error
}

func (dao *RoleDao) Delete(ctx context.Context, id int64) error {
	return dao.db.WithContext(ctx).Model(&model.Role{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"state":       model.StateDelete,
			"update_time": time.Now(),
		}).Error
}

func (dao *RoleDao) List(ctx context.Context, options ...store.Option) ([]model.Role, error) {
	roles := make([]model.Role, 0)
	query := dao.db.WithContext(ctx).Model(&model.Role{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (dao *RoleDao) Count(ctx context.Context, options ...store.Option) (int64, error) {
	var count int64
	query := dao.db.WithContext(ctx).Model(&model.Role{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (dao *RoleDao) ListUserRoles(ctx context.Context, userId int64, options ...store.Option) ([]model.Role, error) {
	roles := make([]model.Role, 0)
	query := dao.db.WithContext(ctx).Model(&model.Role{}).
		Joins("inner join sys_user_role on sys_user_role.role_id = sys_role.id").
		Where("sys_user_role.user_id = ?", userId)

	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
