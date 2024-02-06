// service/imageservice.go
package service

import (
	"myapp/dao"
	"myapp/models"
	_ "myapp/models"
)

type ImageService struct {
	imageDao *dao.ImageDao
}

func NewImageService(imageDao *dao.ImageDao) *ImageService {
	return &ImageService{
		imageDao: imageDao,
	}
}

// DeleteImage 删除指定ID的图片
func (s *ImageService) DeleteImage(imageID uint) error {
	return s.imageDao.DeleteImage(imageID)
}

func (s *ImageService) GetImageByID(id string) (*models.Image, error) {
	return s.imageDao.GetImageByID(id)
}

func (s *ImageService) LikeImage(id string) error {
	// 简化逻辑：每次调用增加 1 个赞
	return s.imageDao.UpdateLikes(id, 1)
}

func (s *ImageService) GetImageFilePathByID(id string) (string, error) {
	return s.imageDao.GetImageFilePathByID(id)
}

func (s *ImageService) CommentOnImage(imageID string, text string) error {
	comment := &models.Comment{
		ImageID: imageID,
		Text:    text,
	}
	return s.imageDao.AddComment(comment)
}
