package middleware

import (
	log "log/slog"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error("panic error %v", err)
				log.Error(string(debug.Stack()))
				ReturnError(c, err.(error))
			}
		}()
		c.Next()
	}
}
