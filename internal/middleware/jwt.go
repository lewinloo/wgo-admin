package middleware

import (
	"gin_template/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int = 200
		var message string
		token := c.GetHeader("Authorization")
		claims, err := utils.ParseToken(token)
		if token == "" || err != nil {
			code = 400
			message = "用户认证失败"
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = 401
			message = "用户认证信息过期"
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": message,
				"success": false,
			})
			c.Abort()
			return
		}

		c.Set("uid", claims.Id)
		c.Set("uuid", claims.UUID)
		c.Next()
	}
}
