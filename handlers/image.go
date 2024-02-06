package handlers

import (
	"github.com/gin-gonic/gin"
	"myapp/service" // 替换为实际的包路径
	"net/http"
	"path/filepath"
)

// 假设这是全局变量，已经在main.go或其他地方初始化
var imageService *service.ImageService

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 生成文件保存路径
	savePath := filepath.Join("uploads", file.Filename)

	// 保存文件
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 这里添加逻辑将文件信息保存到数据库
	// imageService.AddImage(...) // 需要实现这个方法

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully.", "path": savePath})
}

func ShowImage(c *gin.Context) {
	imageID := c.Param("id")
	image, err := imageService.GetImageByID(imageID) // 确保这个方法实现了
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image": image})
}

func LikeImage(c *gin.Context) {
	imageID := c.Param("id")
	if err := imageService.LikeImage(imageID); err != nil { // 确保这个方法实现了
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image liked successfully"})
}

func DownloadImage(c *gin.Context) {
	imageID := c.Param("id")
	filePath, err := imageService.GetImageFilePathByID(imageID) // 确保这个方法实现了
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.File(filePath)
}

func CommentOnImage(c *gin.Context) {
	imageID := c.Param("id")
	commentText := c.PostForm("comment")
	if err := imageService.CommentOnImage(imageID, commentText); err != nil { // 确保这个方法实现了
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}
