package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/data-space/internal/pkg/retcode"
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
			ReturnError(c, retcode.RequestTokenEmpty)
		}
		token := strings.Split(auth, " ")
		if len(token) != 2 || token[0] != "bearer" {
			ReturnError(c, retcode.RequestTokenEmpty)
		}
		claims, err := jwt.ParseToken(token[1])
		if err != nil {
			ReturnError(c, retcode.RequestTokenAuthFail)
		} else {
			c.Set("username", claims.UserName)
			c.Set("userid", claims.UserId)
			c.Next()
		}
	}
}
