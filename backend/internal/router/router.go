package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/hexiaopi/data-space/internal/config"
	"github.com/hexiaopi/data-space/internal/middleware"
	"github.com/hexiaopi/data-space/internal/router/api"
	v1 "github.com/hexiaopi/data-space/internal/router/api/v1"
	"github.com/hexiaopi/data-space/internal/service"
	dao "github.com/hexiaopi/data-space/internal/store/mysql"

	"github.com/hexiaopi/data-space/pkg/metrics"
	"github.com/hexiaopi/data-space/pkg/retcode"
)

type Response struct {
	*retcode.RetCode
	Data interface{} `json:"data,omitempty"`
}

type HandlerFunc func(ctx *gin.Context) (interface{}, error)

func Wrap(handler HandlerFunc) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		response := Response{RetCode: retcode.Success}
		res, err := handler(ctx)
		if err != nil {
			var code *retcode.RetCode
			if errors.As(err, &code) {
				response.RetCode = code
				ctx.Error(code)
			} else {
				response.RetCode = retcode.InternalError
			}
		} else {
			response.Data = res
		}
		metrics.RequestRetcodeCounter.WithLabelValues(path, method, response.Code, response.Desc).Inc()
		ctx.JSON(http.StatusOK, response)
	}
}

func InitRouter() *gin.Engine {
	storeIns := dao.NewDao(config.DBEngine)
	optionIns := dao.NewOption()

	router := gin.New()
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// profiling
	router.GET("/debug/pprof/*any", gin.WrapH(http.DefaultServeMux))
	// metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.NoMethod(Wrap(api.MethodNotAllow))
	router.NoRoute(Wrap(api.PathNotFound))

	apiRouter := router.Group("/api")
	apiRouter.Use(middleware.Recovery(config.Logger))
	apiRouter.Use(middleware.Metrics())
	apiRouter.Use(middleware.Cors())
	apiRouter.Use(middleware.Logger(config.Logger))
	apiRouter.Use(middleware.JWT(config.JWT, middleware.PathContainSkipper("login")))

	v1Router := apiRouter.Group("/v1")

	srv := service.NewService(storeIns, optionIns, config.JWT, config.Logger)

	{
		aclController := v1.NewAclController(srv)
		aclRouter := v1Router.Group("/acl")
		aclRouter.POST("/login", Wrap(aclController.Login))
		aclRouter.POST("/logout", Wrap(aclController.Logout))
		aclRouter.GET("/tree", Wrap(aclController.Tree))
		aclRouter.GET("/user", Wrap(aclController.Info))
	}

	{
		userController := v1.NewUserController(srv)
		userRouter := v1Router.Group("/users")
		userRouter.GET("", Wrap(userController.List))
		userRouter.POST("", Wrap(userController.Create))
		userRouter.GET("/:id", Wrap(userController.Info))
		userRouter.PUT("/:id", Wrap(userController.Update))
		userRouter.DELETE("/:id", Wrap(userController.Delete))
	}

	{
		roleController := v1.NewRoleController(srv)
		roleRouter := v1Router.Group("/roles")
		roleRouter.GET("", Wrap(roleController.List))
		roleRouter.POST("", Wrap(roleController.Create))
		roleRouter.GET("/:id", Wrap(roleController.Get))
		roleRouter.PUT("/:id", Wrap(roleController.Update))
		roleRouter.DELETE("/:id", Wrap(roleController.Delete))
		roleRouter.GET("/:id/menu", Wrap(roleController.GetMenu))
		roleRouter.PUT("/:id/menu", Wrap(roleController.UpdateMenu))
	}

	{
		menuController := v1.NewMenuController(srv)
		menuRouter := v1Router.Group("/menus")
		menuRouter.POST("", Wrap(menuController.Create))
		menuRouter.GET("/tree", Wrap(menuController.Tree))
		menuRouter.PUT("/:id", Wrap(menuController.Update))
		menuRouter.DELETE("/:id", Wrap(menuController.Delete))
	}

	{
		departmentController := v1.NewDepartMentController(srv)
		departmentRouter := v1Router.Group("/departments")
		departmentRouter.GET("", Wrap(departmentController.List))
		departmentRouter.POST("", Wrap(departmentController.Create))
		departmentRouter.PUT("/:id", Wrap(departmentController.Update))
		departmentRouter.DELETE("/:id", Wrap(departmentController.Delete))
	}
	return router
}
