package dao

import (
	"gorm.io/gorm"
	"myapp/global"
	"myapp/models"
)

type ImageDao struct {
	db *gorm.DB
}

func NewImageDao(db *gorm.DB) *ImageDao {
	return &ImageDao{db: db}
}

// 删除图片
func (dao *ImageDao) DeleteImage(imageID uint) error {
	result := global.DB.Delete(&models.Image{}, imageID)
	return result.Error
}
func (dao *ImageDao) GetImageByID(id string) (*models.Image, error) {
	var image models.Image
	err := global.DB.Preload("Comments").First(&image, id).Error
	return &image, err
}

func (dao *ImageDao) UpdateLikes(id string, likes int) error {
	return global.DB.Model(&models.Image{}).Where("id = ?", id).Update("likes", gorm.Expr("likes + ?", likes)).Error
}

func (dao *ImageDao) GetImageFilePathByID(id string) (string, error) {
	var image models.Image
	err := global.DB.Select("path").First(&image, id).Error
	if err != nil {
		return "", err
	}
	return image.Path, nil
}

func (dao *ImageDao) AddComment(comment *models.Comment) error {
	return global.DB.Create(comment).Error
}
func (dao *ImageDao) GetAllImages(image *[]*models.Image) error {
	err := global.DB.Find(&image)
	return err.Error

}
