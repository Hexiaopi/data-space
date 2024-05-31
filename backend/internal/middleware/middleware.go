package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/pkg/retcode"
)

type SkipperFunc func(ctx *gin.Context) bool

func PathPrefixSkipper(prefixs ...string) SkipperFunc {
	return func(ctx *gin.Context) bool {
		path := ctx.Request.URL.Path
		pathLength := len(path)
		for _, prefix := range prefixs {
			pl := len(prefix)
			if pathLength >= pl && path[:pl] == prefix {
				return true
			}
		}
		return false
	}
}

func PathContainSkipper(prefixs ...string) SkipperFunc {
	return func(ctx *gin.Context) bool {
		path := ctx.Request.URL.Path
		for _, prefix := range prefixs {
			if strings.Contains(path, prefix) {
				return true
			}
		}
		return false
	}
}

func SkipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}

func ReturnError(ctx *gin.Context, err error) {
	path := ctx.Request.URL.Path
	method := ctx.Request.Method
	var code *retcode.RetCode
	if !errors.As(err, &code) {
		code = retcode.UnknownError
	}
	retcode.RequestRetcodeCounter.WithLabelValues(path, method, code.Code, code.Desc).Inc()
	ctx.Abort()
}
