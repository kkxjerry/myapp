package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowUploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{
		"title": "Login Page",
	})
}
