package v1

import (
	"strconv"

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
	userId := ctx.Value(global.UserId).(int64)
	res, err := c.srv.Users().Info(ctx.Request.Context(), userId)
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
	err := c.srv.Users().Create(ctx.Request.Context(), &req)
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
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return nil, global.RequestIllegal
	}
	req.Id = id
	err = c.srv.Users().Update(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *UserController) Delete(ctx *gin.Context) (interface{}, error) {
	var req service.DeleteUserRequest
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return nil, global.RequestIllegal
	}
	req.Id = id
	err = c.srv.Users().Delete(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *UserController) List(ctx *gin.Context) (interface{}, error) {
	var req service.ListUserRequest
	departmentId, err := strconv.ParseInt(ctx.Query("department_id"), 10, 64)
	if err != nil {
		return nil, global.RequestIllegal
	}
	req.Name = ctx.Query("name")
	state, _ := strconv.Atoi(ctx.Query("state"))
	req.State = uint8(state)
	req.DepartmentId = departmentId
	pageNum := ctx.Query("page_num")
	req.PageNum, _ = strconv.Atoi(pageNum)
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	pageSize := ctx.Query("page_size")
	req.PageSize, _ = strconv.Atoi(pageSize)

	res, err := c.srv.Users().List(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
