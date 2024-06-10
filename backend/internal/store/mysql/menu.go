package mysql

import (
	"context"

	"gorm.io/gorm"

	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
)

type MenuDao struct {
	db *gorm.DB
}

func NewMenuDao(db *gorm.DB) *MenuDao {
	return &MenuDao{db: db}
}

func (dao *MenuDao) Get(ctx context.Context, options ...store.Option) (*model.Menu, error) {
	menu := model.Menu{}
	query := dao.db.WithContext(ctx).Model(&model.Menu{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.First(&menu).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (dao *MenuDao) Create(ctx context.Context, menu *model.Menu) error {
	return dao.db.WithContext(ctx).Create(menu).Error
}

func (dao *MenuDao) Update(ctx context.Context, menu *model.Menu) error {
	return dao.db.WithContext(ctx).Save(menu).Error
}

func (dao *MenuDao) Delete(ctx context.Context, id int64) error {
	menu := model.Menu{ID: id}
	return dao.db.WithContext(ctx).Delete(&menu).Error
}

func (dao *MenuDao) List(ctx context.Context, options ...store.Option) ([]model.Menu, error) {
	menus := make([]model.Menu, 0)
	query := dao.db.WithContext(ctx).Model(&model.Menu{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (dao *MenuDao) Count(ctx context.Context, options ...store.Option) (int64, error) {
	var count int64
	query := dao.db.WithContext(ctx).Model(&model.Menu{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (dao *MenuDao) ListRoleMenus(ctx context.Context, roleId int64, options ...store.Option) ([]model.Menu, error) {
	menus := make([]model.Menu, 0)
	query := dao.db.WithContext(ctx).Model(&model.Menu{}).
		Joins("inner join sys_menu_role on sys_menu_role.menu_id=sys_menu.id").
		Where("sys_menu_role.role_id=?", roleId).
		Where("sys_menu.state=?", model.StatusEnable).
		Order("sys_menu.order asc")
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}
