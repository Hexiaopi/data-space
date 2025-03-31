package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/global"
)

func PathNotFound(ctx *gin.Context) (interface{}, error) {
	log.Println(ctx.Request.URL.Path)
	return nil, global.RequestPathNotFound
}

func MethodNotAllow(ctx *gin.Context) (interface{}, error) {
	return nil, global.RequestMethodNotAllow
}
