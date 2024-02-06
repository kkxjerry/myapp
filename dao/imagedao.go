package dao

import (
	"gorm.io/gorm"
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
	result := dao.db.Delete(&models.Image{}, imageID)
	return result.Error
}
func (dao *ImageDao) GetImageByID(id uint) (*models.Image, error) {
	var image models.Image
	err := dao.db.Preload("Comments").First(&image, id).Error
	return &image, err
}

func (dao *ImageDao) UpdateLikes(id uint, likes int) error {
	return dao.db.Model(&models.Image{}).Where("id = ?", id).Update("likes", gorm.Expr("likes + ?", likes)).Error
}

func (dao *ImageDao) GetImageFilePathByID(id uint) (string, error) {
	var image models.Image
	err := dao.db.Select("path").First(&image, id).Error
	if err != nil {
		return "", err
	}
	return image.Path, nil
}

func (dao *ImageDao) AddComment(comment *models.Comment) error {
	return dao.db.Create(comment).Error
}
