package handlers

import (
	"github.com/gin-gonic/gin"
	"myapp/models"
	"net/http"
)

func GetImages(c *gin.Context) {
	images := make([]*models.Image, 0)
	err := imageService.GetAllImages(&images) // 假设这个函数返回所有图片的实例列表
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	// 仅提取图片路径
	var paths []string
	for _, img := range images {
		paths = append(paths, img.Path)
	}

	c.JSON(http.StatusOK, gin.H{"images": paths})
}
