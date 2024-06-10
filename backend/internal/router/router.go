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
	apiRouter.Use(middleware.Metrics())
	apiRouter.Use(middleware.Cors())
	apiRouter.Use(middleware.JWT(config.JWT, middleware.PathContainSkipper("login")))

	v1Router := apiRouter.Group("/v1")

	aclRouter := v1Router.Group("/acl")

	srv := service.NewService(storeIns, optionIns, config.JWT)

	userController := v1.NewUserController(srv)
	menuController := v1.NewMenuController(srv)

	aclRouter.POST("/login", Wrap(userController.Login))
	aclRouter.POST("/logout", Wrap(userController.Logout))
	aclRouter.GET("/tree", Wrap(menuController.Tree))

	v1Router.GET("/users", Wrap(userController.List))
	v1Router.GET("/user", Wrap(userController.Info))
	v1Router.POST("/user", Wrap(userController.Create))
	v1Router.PUT("/user", Wrap(userController.Update))
	v1Router.DELETE("/user", Wrap(userController.Delete))

	return router
}
