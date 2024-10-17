package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
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

// @Summary 登录接口
// @Description 用户登录生成Token
// @Tags acl
// @Produce json
// @Accept json
// @param LoginRequest body service.LoginRequest true "用户信息"
// @Success 200 {object} router.Response{data=service.LoginResponse} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/acl/login [post]
func (c *AclController) Login(ctx *gin.Context) (interface{}, error) {
	var req service.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, global.RequestUnMarshalError
	}
	remoteIp := ctx.ClientIP()
	userAgent := ctx.Request.UserAgent()
	loginLog := model.LoginLog{UserName: req.UserName, RemoteIP: remoteIp, UserAgent: userAgent, LoginResult: 1}
	res, err := c.srv.Users().Login(ctx.Request.Context(), &req)
	if err != nil {
		loginLog.LoginResult = 2
		loginLog.ResultDetail = err.Error()
		_ = c.srv.LoginLogs().Create(ctx.Request.Context(), &loginLog)
		return nil, err
	}
	_ = c.srv.LoginLogs().Create(ctx.Request.Context(), &loginLog)
	return res, nil
}

// @Summary 登出接口
// @Description 当前登录用户登出
// @Tags acl
// @Produce json
// @Accept json
// @Security JWT
// @Success 200 {object} router.Response{} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/acl/logout [post]
func (c *AclController) Logout(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}

// @Summary 访问登录信息
// @Description 当前登录用户的相关信息
// @Tags acl
// @Produce json
// @Accept json
// @Security JWT
// @Success 200 {object} router.Response{data=entity.UserInfo} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/acl/user [get]
func (c *AclController) User(ctx *gin.Context) (interface{}, error) {
	userId := ctx.Value(global.UserId).(int64)
	res, err := c.srv.Users().Info(ctx.Request.Context(), userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// @Summary 访问菜单树
// @Description 访问当前登录用户的菜单树
// @Tags acl
// @Produce json
// @Accept json
// @Security JWT
// @Success 200 {object} router.Response{data=entity.MenuTree} "成功"
// @Failure 200 {object} router.Response{} "失败"
// @Router /api/v1/acl/tree [get]
func (c *AclController) Tree(ctx *gin.Context) (interface{}, error) {
	userId := ctx.Value(global.UserId).(int64)
	var req = service.AclMenuTreeRequest{UserId: userId}
	return c.srv.Menus().AclTree(ctx.Request.Context(), &req)
}
