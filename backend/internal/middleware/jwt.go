package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/pkg/jwt"
)

// JWT 身份验证
func JWT(jwt jwt.JWT, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			ReturnError(c, global.RequestTokenEmpty)
			return
		}
		token := strings.Split(auth, " ")
		if len(token) != 2 || token[0] != "Bearer" {
			ReturnError(c, global.RequestTokenEmpty)
			return
		}
		claims, err := jwt.ParseToken(token[1])
		if err != nil {
			log.Println(err)
			ReturnError(c, global.RequestTokenAuthFail)
			return
		} else {
			c.Set(global.UserName, claims.UserName)
			c.Set(global.UserId, claims.UserId)
			c.Next()
		}
	}
}
