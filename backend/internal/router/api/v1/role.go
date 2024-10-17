package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/service"
)

type RoleController struct {
	srv service.Service
}

func NewRoleController(srv service.Service) *RoleController {
	return &RoleController{
		srv: srv,
	}
}

// @Summary 角色列表接口
// @Description 查询角色列表
// @Tags role
// @Produce json
// @Accept json
// @Security JWT
// @param RoleListRequest query service.RoleListRequest true "请求参数"
// @Success 200 {object} router.Response{data=service.RoleListResponse} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/roles [get]
func (c *RoleController) List(ctx *gin.Context) (interface{}, error) {
	var req service.RoleListRequest
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
	name := ctx.Query("name")
	if name != "" {
		req.Name = name
	}
	state := ctx.Query("state")
	if state != "" {
		s, _ := strconv.ParseUint(state, 10, 64)
		if s > 0 {
			req.State = uint8(s)
		}
	}
	res, err := c.srv.Roles().List(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// @Summary 角色创建接口
// @Description 创建角色信息
// @Tags role
// @Produce json
// @Accept json
// @Security JWT
// @param RoleCreateRequest body service.RoleCreateRequest false "请求参数"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/roles [post]
func (c *RoleController) Create(ctx *gin.Context) (interface{}, error) {
	var req service.RoleCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	err := c.srv.Roles().Create(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// @Summary 角色修改接口
// @Description 根据角色id修改角色信息
// @Tags role
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "角色id"
// @param RoleUpdateRequest body service.RoleUpdateRequest false "请求参数"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/roles/{id} [put]
func (c *RoleController) Update(ctx *gin.Context) (interface{}, error) {
	var req service.RoleUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Roles().Update(ctx.Request.Context(), id, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// @Summary 角色删除接口
// @Description 根据角色id删除角色信息
// @Tags role
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "角色id"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/roles/{id} [delete]
func (c *RoleController) Delete(ctx *gin.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Roles().Delete(ctx.Request.Context(), id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// @Summary 角色详情接口
// @Description 根据角色id查询角色信息
// @Tags role
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "角色id"
// @Success 200 {object} router.Response{data=service.RoleInfoResponse} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/roles/{id} [get]
func (c *RoleController) Get(ctx *gin.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	res, err := c.srv.Roles().Get(ctx.Request.Context(), id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// @Summary 角色菜单接口
// @Description 根据角色id查询菜单信息
// @Tags role
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "角色ID"
// @Success 200 {object} router.Response{data=entity.MenuTree} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/roles/{id}/menu [get]
func (c *RoleController) GetMenu(ctx *gin.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	res, err := c.srv.Roles().GetMenu(ctx.Request.Context(), id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// @Summary 修改角色菜单接口
// @Description 修改角色对应的菜单信息
// @Tags role
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "角色ID"
// @param UpdateMenuRequest body service.UpdateMenuRequest false "请求参数"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/roles/{id}/menu [put]
func (c *RoleController) UpdateMenu(ctx *gin.Context) (interface{}, error) {
	var req service.UpdateMenuRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := c.srv.Roles().UpdateMenu(ctx.Request.Context(), id, &req); err != nil {
		return nil, err
	}
	return nil, nil
}
