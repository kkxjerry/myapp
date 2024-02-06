// service/imageservice.go
package service

import (
	"myapp/dao"
	"myapp/global"
	"myapp/models"
	_ "myapp/models"
	"strconv"
)

var imageService *ImageService

type ImageService struct {
	imageDao *dao.ImageDao
}

func init() {
	imageDao := dao.NewImageDao(global.DB)
	imageService = NewImageService(imageDao)

}

func NewImageService(imageDao *dao.ImageDao) *ImageService {
	return &ImageService{
		imageDao: imageDao,
	}
}

// DeleteImage 删除指定ID的图片
func (s *ImageService) DeleteImage(imageID uint) error {
	return imageService.imageDao.DeleteImage(imageID)
}

func (imageService *ImageService) GetImageByID(id string) (*models.Image, error) {
	return imageService.imageDao.GetImageByID(id)
}

func (s *ImageService) LikeImage(id string) error {
	// 简化逻辑：每次调用增加 1 个赞
	return imageService.imageDao.UpdateLikes(id, 1)
}

func (s *ImageService) GetImageFilePathByID(id string) (string, error) {
	return imageService.imageDao.GetImageFilePathByID(id)
}

func (s *ImageService) CommentOnImage(imageID string, text string) error {
	id, err := strconv.Atoi(imageID)
	if err != nil {
	}
	imageid := uint(id)
	comment := &models.Comment{
		ImageID: imageid,
		Text:    text,
	}
	return imageService.imageDao.AddComment(comment)
}

func (s *ImageService) GetAllImages(image *[]*models.Image) error {
	err := imageService.imageDao.GetAllImages(image)
	return err
}
