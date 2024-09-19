package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/service"
)

type DepartmentController struct {
	srv service.Service
}

func NewDepartMentController(srv service.Service) *DepartmentController {
	return &DepartmentController{
		srv: srv,
	}
}

func (c *DepartmentController) List(ctx *gin.Context) (interface{}, error) {
	var req service.DepartmentListRequest
	pageNum := ctx.Query("page_num")
	req.PageNum, _ = strconv.Atoi(pageNum)

	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	pageSize := ctx.Query("page_size")
	req.PageSize, _ = strconv.Atoi(pageSize)
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	res, err := c.srv.Departments().List(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *DepartmentController) Create(ctx *gin.Context) (interface{}, error) {
	var req service.DepartmentCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	err := c.srv.Departments().Create(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *DepartmentController) Update(ctx *gin.Context) (interface{}, error) {
	var req service.DepartmentUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Departments().Update(ctx.Request.Context(), id, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *DepartmentController) Delete(ctx *gin.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Departments().Delete(ctx.Request.Context(), id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
