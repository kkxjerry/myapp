package main

import (
	"github.com/gin-gonic/gin"
	"myapp/global"
	"myapp/handlers"
	middleware "myapp/middlware"
	"net/http"
	// 引入其他必要的包
)

func main() {
	global.LoadConfig()
	global.InitDB()
	//userDao := dao.NewUserDao(global.DB)
	//userService := service.NewUserService(userDao)

	r := gin.Default()

	// 日志中间件
	r.Use(middleware.LoggingMiddleware())
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 公开路由
	r.GET("/login", handlers.ShowLoginPage)
	r.POST("/login", handlers.Login)
	r.GET("/image/:id", handlers.ShowImage)
	r.POST("/image/:id/like", handlers.LikeImage)
	r.POST("/image/:id/comment", handlers.CommentOnImage)

	// 认证中间件
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
		auth.GET("/images", handlers.GetImages)
		auth.POST("/logout", handlers.Logout)
		auth.GET("/upload", handlers.ShowUploadPage)
		auth.POST("/upload", handlers.UploadImage)
		auth.GET("/image/:id/download", handlers.DownloadImage)
	}
	r.GET("/admin/login", handlers.ShowAdminLoginPage)
	r.POST("/admin/login", handlers.AdminLogin)
	// 管理员路由
	admin := r.Group("/admin")

	admin.Use(middleware.AdminAuthMiddleware())
	{
		admin.GET("/export_users", handlers.ExportUsers)
		admin.DELETE("/image/:id", handlers.DeleteImage)
	}

	r.Run(global.Config.Server.Port)
}
