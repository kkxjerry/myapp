// handlers/admin.go
package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	// 引入 models 和其他必要的包
)

func ExportUsers(c *gin.Context) {
	filename := "users.csv"
	if err := userService.ExportUsersCSV(filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(filename)
}
func DeleteImage(c *gin.Context) {
	// 从URL参数中获取imageID
	imageID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}
	// 调用服务层的DeleteImage
	if err := imageService.DeleteImage(uint(imageID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}

func AdminLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证凭据；这是一个占位符
	if username == "admin" && password == "secret" {
		// 设置会话或 Cookie 以指示成功登录
		c.SetCookie("admin", "true", 3600, "/", "", false, true)
		c.Redirect(http.StatusSeeOther, "/admin")
	} else {
		c.HTML(http.StatusUnauthorized, "admin_login.html", gin.H{"Error": "无效的凭据"})
	}
}
