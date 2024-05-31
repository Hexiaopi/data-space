package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hexiaopi/data-space/internal/pkg/retcode"
	"github.com/hexiaopi/data-space/internal/service"
	"github.com/hexiaopi/data-space/internal/store"
	"github.com/hexiaopi/data-space/pkg/jwt"
)

type UserController struct {
	srv service.Service
}

func NewUserController(store store.Factory, option store.Option, jwt jwt.JWT) *UserController {
	return &UserController{
		srv: service.NewService(store, option, jwt),
	}
}

func (c *UserController) Login(ctx *gin.Context) (interface{}, error) {
	var req service.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, retcode.RequestUnMarshalError
	}
	res, err := c.srv.Users().Login(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
