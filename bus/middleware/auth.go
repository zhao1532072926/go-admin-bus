package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin-bus/bus/utils"
	"net/http"
	"strings"
)

func JWTAuth(jwt *utils.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限验证失败",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限验证失败",
			})
			c.Abort()
			return
		}

		token, err := jwt.ValidateToken(parts[1])
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限验证失败",
			})
			c.Abort()
			return
		}

		c.Next()
	}
} 