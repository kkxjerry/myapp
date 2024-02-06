package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowAdminLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_login.html", nil)
}
