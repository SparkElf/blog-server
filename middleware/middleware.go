package middleware

import (
	"strings"

	"github.com/SparkElf/blog-server/middleware/jwt"
	. "github.com/SparkElf/blog-server/utils/exception"
	"github.com/gin-gonic/gin"
)

// jwt 中间件
/**
 * 接受到的token示例:
 * Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZvcmVzdGxpZ2h0ZWxmIiwicGFzc3dvcmQiOiIxMjM0NTYiLCJleHAiOjE2MzA4NDMxNjUsImlzcyI6ImdpbmJsb2cifQ._7EGKcRDoo_E2iedRCMPcalpEqp5ZpQl3HEypt4GrR0
 */
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		println("token:" + authHeader)

		var claims *jwt.Claim
		var e Err
		if authHeader == "" {
			e.Code = ERROR_TOKEN_NOT_EXIST
		} else {
			token := strings.SplitN(authHeader, " ", 2)
			if len(token) != 2 || token[0] != "Bearer" {
				e.Code = ERROR_TOKEN_TYPE_WRONG
			} else {
				claims, e = jwt.CheckToken(token[1])
				println("code:" + e.Error())
			}

		}

		if e.Code != SUCCESE {
			c.Abort()
		} else {
			c.Set("username", claims.Username)
		}
		c.Next()
	}
}
