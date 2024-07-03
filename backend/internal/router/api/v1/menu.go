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

func (c *MenuController) Tree(ctx *gin.Context) (interface{}, error) {
	var req = service.MenuTreeRequest{}
	name := ctx.Query("name")
	if name != "" {
		req.Name = name
	}
	return c.srv.Menus().Tree(ctx, &req)
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
	res, err := c.srv.Menus().List(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *MenuController) Create(ctx *gin.Context) (interface{}, error) {
	var req service.MenuCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	err := c.srv.Menus().Create(ctx, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *MenuController) Update(ctx *gin.Context) (interface{}, error) {
	var req service.MenuUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Menus().Update(ctx, id, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *MenuController) Delete(ctx *gin.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Menus().Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
