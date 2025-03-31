package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hexiaopi/data-space/internal/service"
)

type LoginLogController struct {
	srv service.Service
}

func NewLoginLogController(srv service.Service) *LoginLogController {
	return &LoginLogController{
		srv: srv,
	}
}

// @Summary 登录日志查询接口
// @Description 查询登录日志信息
// @Tags log
// @Produce json
// @Accept json
// @Security JWT
// @param LoginLogListRequest query service.LoginLogListRequest true "请求参数"
// @Success 200 {object} router.Response{data=service.LoginLogListResponse} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/log/logins [get]
func (c *LoginLogController) List(ctx *gin.Context) (interface{}, error) {
	var req service.LoginLogListRequest
	req.UserName = ctx.Query("user_name")
	req.LoginResult, _ = strconv.Atoi(ctx.Query("login_result"))
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
	res, err := c.srv.LoginLogs().List(ctx.Request.Context(), &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
