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
	"github.com/hexiaopi/data-space/internal/pkg/retcode"
	"github.com/hexiaopi/data-space/internal/router/api"
	v1 "github.com/hexiaopi/data-space/internal/router/api/v1"
	dao "github.com/hexiaopi/data-space/internal/store/mysql"
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
			}
		} else {
			response.Data = res
		}
		retcode.RequestRetcodeCounter.WithLabelValues(path, method, response.Code, response.Desc).Inc()
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

	userRouter := v1Router.Group("/user")
	userController := v1.NewUserController(storeIns, optionIns, config.JWT)
	userRouter.POST("/login", Wrap(userController.Login))

	return router
}
