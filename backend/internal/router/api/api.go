package api

import (
	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/pkg/retcode"
)

func PathNotFound(ctx *gin.Context) (interface{}, error) {
	return nil, retcode.RequestPathNotFound
}

func MethodNotAllow(ctx *gin.Context) (interface{}, error) {
	return nil, retcode.RequestMethodNotAllow
}
