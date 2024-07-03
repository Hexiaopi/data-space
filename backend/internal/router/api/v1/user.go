package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/service"
)

type UserController struct {
	srv service.Service
}

func NewUserController(srv service.Service) *UserController {
	return &UserController{
		srv: srv,
	}
}

func (c *UserController) Info(ctx *gin.Context) (interface{}, error) {
	res, err := c.srv.Users().Info(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *UserController) Create(ctx *gin.Context) (interface{}, error) {
	var req service.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	err := c.srv.Users().Create(ctx, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *UserController) Update(ctx *gin.Context) (interface{}, error) {
	var req service.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	err := c.srv.Users().Update(ctx, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *UserController) Delete(ctx *gin.Context) (interface{}, error) {
	var req service.DeleteUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	err := c.srv.Users().Delete(ctx, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *UserController) List(ctx *gin.Context) (interface{}, error) {
	var req service.ListUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	res, err := c.srv.Users().List(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
