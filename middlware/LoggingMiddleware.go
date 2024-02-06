package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		t := time.Now()

		// 设置请求的例子变量
		c.Set("example", "12345")

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的状态码
		status := c.Writer.Status()
		log.Printf("Status: %d", status)
	}
}
