package mysql

import (
	"context"
	"time"

	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
	"gorm.io/gorm"
)

type DepartmentDao struct {
	db *gorm.DB
}

func NewDepartmentDao(db *gorm.DB) *DepartmentDao {
	return &DepartmentDao{db: db}
}

func (dao *DepartmentDao) List(ctx context.Context, options ...store.Option) ([]model.Department, error) {
	menus := make([]model.Department, 0)
	query := dao.db.WithContext(ctx).Model(&model.Department{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (dao *DepartmentDao) Count(ctx context.Context, options ...store.Option) (int64, error) {
	var count int64
	query := dao.db.WithContext(ctx).Model(&model.Department{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (dao *DepartmentDao) Get(ctx context.Context, options ...store.Option) (*model.Department, error) {
	department := model.Department{}
	query := dao.db.WithContext(ctx).Model(&model.Department{})
	for _, option := range options {
		option.(Option)(query)
	}
	if err := query.First(&department).Error; err != nil {
		return nil, err
	}
	return &department, nil
}

func (dao *DepartmentDao) Create(ctx context.Context, department *model.Department) error {
	if department == nil {
		return nil
	}
	now := time.Now()
	department.CreateTime = now
	department.UpdateTime = now
	return dao.db.WithContext(ctx).Create(department).Error
}

func (dao *DepartmentDao) Update(ctx context.Context, department *model.Department) error {
	return dao.db.WithContext(ctx).Model(&model.Department{}).
		Where("id = ?", department.ID).
		Updates(map[string]interface{}{
			"name":        department.Name,
			"desc":        department.Desc,
			"state":       department.State,
			"update_time": time.Now(),
		}).Error
}

func (dao *DepartmentDao) Delete(ctx context.Context, id int64) error {
	return dao.db.WithContext(ctx).Model(&model.Department{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"state":       model.StateDelete,
			"update_time": time.Now(),
		}).Error
}
