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
	userId := ctx.Value(global.UserId).(int64)
	var req = service.MenuTreeRequest{UserId: userId}
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
