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
	res, err := c.srv.Roles().List(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *RoleController) Create(ctx *gin.Context) (interface{}, error) {
	var req service.RoleCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	err := c.srv.Roles().Create(ctx, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *RoleController) Update(ctx *gin.Context) (interface{}, error) {
	var req service.RoleUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Roles().Update(ctx, id, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *RoleController) Delete(ctx *gin.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	err := c.srv.Roles().Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *RoleController) Get(ctx *gin.Context) (interface{}, error) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	res, err := c.srv.Roles().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
