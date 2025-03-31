package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/service"
)

type MenuController struct {
	srv service.Service
}

func NewMenuController(srv service.Service) *MenuController {
	return &MenuController{
		srv: srv,
	}
}

// @Summary 菜单树接口
// @Description 查询菜单树
// @Tags menu
// @Produce json
// @Accept json
// @Security JWT
// @param name query string false "名称"
// @Success 200 {object} router.Response{data=entity.MenuTree} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/menus/tree [get]
func (c *MenuController) Tree(ctx *gin.Context) (interface{}, error) {
	var req = service.MenuTreeRequest{}
	name := ctx.Query("name")
	if name != "" {
		req.Name = name
	}
	return c.srv.Menus().Tree(ctx.Request.Context(), &req)
}

func (c *MenuController) List(ctx *gin.Context) (interface{}, error) {
	var req service.MenuListRequest
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
	res, err := c.srv.Menus().List(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// @Summary 创建菜单接口
// @Description 创建菜单信息
// @Tags menu
// @Produce json
// @Accept json
// @Security JWT
// @param MenuCreateRequest body service.MenuCreateRequest true "请求参数"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/menus [post]
func (c *MenuController) Create(ctx *gin.Context) (interface{}, error) {
	var req service.MenuCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	err := c.srv.Menus().Create(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// @Summary 修改菜单接口
// @Description 修改菜单信息
// @Tags menu
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "菜单id"
// @param MenuUpdateRequest body service.MenuUpdateRequest true "请求参数"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/menus/{id} [put]
func (c *MenuController) Update(ctx *gin.Context) (interface{}, error) {
	var req service.MenuUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Menus().Update(ctx.Request.Context(), id, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// @Summary 删除菜单接口
// @Description 根据id删除菜单
// @Tags menu
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "菜单id"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/menus/{id} [delete]
func (c *MenuController) Delete(ctx *gin.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Menus().Delete(ctx.Request.Context(), id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
