package middleware

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/pkg/logger"
)

func Recovery(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("panic error %v", err)
				log.Error(string(debug.Stack()))
				ReturnError(c, err.(error))
			}
		}()
		c.Next()
	}
}
