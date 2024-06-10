package v1

import (
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
