package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里的逻辑应根据你的应用需求来定制
		// 例如, 检查用户角色是否为 "admin"
		isAdmin, exists := c.Get("is_admin")
		if !exists || !isAdmin.(bool) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
