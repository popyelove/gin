package jwt

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"Gin/common/tools/e"
	"Gin/common/tools/jwt"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token :=c.Request.Header.Get("token")
		if token == "" {
			code = e.ERROR_AUTH_TOKEN_NOT_EXIST
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
