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

// @Summary 部门列表接口
// @Description 查询部门列表
// @Tags department
// @Produce json
// @Accept json
// @Security JWT
// @param request query service.DepartmentListRequest true "请求参数"
// @Success 200 {object} router.Response{data=service.DepartmentListResponse} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/departments [get]
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

// @Summary 创建部门接口
// @Description 创建新的部门
// @Tags department
// @Produce json
// @Accept json
// @Security JWT
// @param DepartmentCreateRequest body service.DepartmentCreateRequest true "请求参数"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/departments [post]
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

// @Summary 修改部门接口
// @Description 根据部门id修改部门信息
// @Tags department
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "部门id"
// @param DepartmentUpdateRequest body service.DepartmentUpdateRequest true "请求参数"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/departments/{id} [put]
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

// @Summary 删除部门接口
// @Description 根据部门id删除部门信息
// @Tags department
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "部门id"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/departments/{id} [delete]
func (c *DepartmentController) Delete(ctx *gin.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Departments().Delete(ctx.Request.Context(), id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
