package v1

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/service"
)

type AclController struct {
	srv service.Service
}

func NewAclController(srv service.Service) *AclController {
	return &AclController{
		srv: srv,
	}
}

func (c *AclController) Login(ctx *gin.Context) (interface{}, error) {
	var req service.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	reqCtx := context.WithValue(ctx.Request.Context(), "remote-ip", ctx.ClientIP())
	reqCtx = context.WithValue(reqCtx, "user-agent", ctx.Request.UserAgent())
	res, err := c.srv.Users().Login(reqCtx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AclController) Logout(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}

func (c *AclController) Info(ctx *gin.Context) (interface{}, error) {
	userId := ctx.Value(global.UserId).(int64)
	res, err := c.srv.Users().Info(ctx.Request.Context(), userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *AclController) Tree(ctx *gin.Context) (interface{}, error) {
	userId := ctx.Value(global.UserId).(int64)
	var req = service.AclMenuTreeRequest{UserId: userId}
	return c.srv.Menus().AclTree(ctx.Request.Context(), &req)
}
