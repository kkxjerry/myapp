package handlers

import (
	"github.com/gin-gonic/gin"
	"myapp/service"
	"net/http"
)

// 依赖注入: 服务层实例
var userService *service.UserService

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, authenticated := userService.Authenticate(username, password)
	if !authenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// 设置登录状态的 Cookie
	c.SetCookie("user_id", string(user.ID), 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Logout(c *gin.Context) {
	// 清除登录状态的 Cookie
	c.SetCookie("user_id", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
