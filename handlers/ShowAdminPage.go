package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowAdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"title": "Admin Page",
	})
}
