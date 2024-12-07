package middleware

import (
	"fmt"
	"go-admin-bus/bus/config"
	"go-admin-bus/bus/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth(jwt *utils.JWT, serviceNames config.ServiceNames) gin.HandlerFunc {
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

		claims, err := jwt.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  fmt.Sprintf("权限验证失败, %v", err),
			})
			c.Abort()
			return
		}

		// 获取 token 中的 service 字段
		service := claims["service"].(string)
		if service != serviceNames.Name {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限验证失败, 用户权限与服务不匹配",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
