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

// @Summary 用户详情接口
// @Description 根据用户id查询用户信息
// @Tags user
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "用户id"
// @Success 200 {object} router.Response{data=entity.UserInfo} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/users/{id} [get]
func (c *UserController) Info(ctx *gin.Context) (interface{}, error) {
	userId := ctx.Value(global.UserId).(int64)
	res, err := c.srv.Users().Info(ctx.Request.Context(), userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// @Summary 创建用户接口
// @Description 创建用户信息
// @Tags user
// @Produce json
// @Accept json
// @Security JWT
// @param CreateUserRequest body service.CreateUserRequest false "请求参数"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/users [post]
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

// @Summary 修改用户接口
// @Description 修改用户信息
// @Tags user
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "用户id"
// @param UpdateUserRequest body service.UpdateUserRequest false "请求参数"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/users/{id} [put]
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

// @Summary 删除用户接口
// @Description 删除用户信息
// @Tags user
// @Produce json
// @Accept json
// @Security JWT
// @param id path int true "用户id"
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/users/{id} [delete]
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

// @Summary 用户列表接口
// @Description 查询用户列表信息
// @Tags user
// @Produce json
// @Accept json
// @Security JWT
// @param request query service.ListUserRequest true "请求参数"
// @Success 200 {object} router.Response{data=service.ListUserResponse} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/users [get]
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
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	res, err := c.srv.Users().List(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
