package service

import (
	"context"

	"github.com/hexiaopi/data-space/internal/entity"
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
	"github.com/hexiaopi/data-space/pkg/logger"
)

type DepartmentSrv interface {
	List(ctx context.Context, param *DepartmentListRequest) (*DepartmentListResponse, error)
	Create(ctx context.Context, req *DepartmentCreateRequest) error
	Update(ctx context.Context, id int64, req *DepartmentUpdateRequest) error
	Delete(ctx context.Context, id int64) error
}

type DepartmentService struct {
	store  store.Factory
	option store.Option
	log    logger.Logger
}

func NewDepartmentService(store store.Factory, option store.Option, log logger.Logger) *DepartmentService {
	return &DepartmentService{
		store:  store,
		option: option,
		log:    log,
	}
}

type DepartmentListRequest struct {
	Name     string `json:"name"`      //名称
	State    uint8  `json:"state"`     //状态
	PageNum  int    `json:"page_num"`  //页码
	PageSize int    `json:"page_size"` //页大小
}

type DepartmentListResponse struct {
	List  []entity.Department `json:"list"`  //列表
	Total int64               `json:"total"` //总数
}

func (svc *DepartmentService) List(ctx context.Context, req *DepartmentListRequest) (*DepartmentListResponse, error) {
	var res DepartmentListResponse
	options := make([]store.Option, 0)
	if req.Name != "" {
		options = append(options, svc.option.WithName(req.Name))
	}
	if req.State > 0 {
		options = append(options, svc.option.WithState(req.State))
	}
	count, err := svc.store.Departments().Count(ctx, options...)
	if err != nil {
		svc.log.Errorf("store department count", err)
		return nil, global.DepartmentCountFail
	}
	if count == 0 {
		return &res, nil
	}
	options = append(options, svc.option.WithPage(req.PageNum, req.PageSize))
	departments, err := svc.store.Departments().List(ctx, options...)
	if err != nil {
		svc.log.Errorf("store department list", err)
		return nil, global.DepartmentListFail
	}
	res.List = entity.ToDepartments(departments)
	res.Total = count
	return &res, nil
}

type DepartmentCreateRequest struct {
	Name  string `json:"name"`  //名称
	Desc  string `json:"desc"`  //描述
	State uint8  `json:"state"` //状态
}

func (svc *DepartmentService) Create(ctx context.Context, req *DepartmentCreateRequest) error {
	var department model.Department
	department.Name = req.Name
	department.Desc = req.Desc
	department.State = req.State
	if err := svc.store.Departments().Create(ctx, &department); err != nil {
		svc.log.Errorf("store department create", err)
		return global.DepartmentCreateFail
	}
	return nil
}

type DepartmentUpdateRequest struct {
	Name  string `json:"name"`  //名称
	Desc  string `json:"desc"`  //描述
	State uint8  `json:"state"` //状态
}

func (svc *DepartmentService) Update(ctx context.Context, id int64, req *DepartmentUpdateRequest) error {
	department, err := svc.store.Departments().Get(ctx, svc.option.WithId(id))
	if err != nil {
		svc.log.Errorf("store department get", err)
		return global.DepartmentGetFail
	}
	department.Name = req.Name
	department.Desc = req.Desc
	department.State = req.State
	if err := svc.store.Departments().Update(ctx, department); err != nil {
		svc.log.Errorf("store department update", err)
		return global.DepartmentUpdateFail
	}
	return nil
}

func (svc *DepartmentService) Delete(ctx context.Context, id int64) error {
	if err := svc.store.Departments().Delete(ctx, id); err != nil {
		svc.log.Errorf("store department delete", err)
		return global.DepartmentDeleteFail
	}
	return nil
}
