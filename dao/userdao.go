package dao

import (
	"gorm.io/gorm"
	"myapp/models"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// 通过用户名查找用户
func (dao *UserDao) FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := dao.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (dao *UserDao) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := dao.db.Find(&users)
	return users, result.Error
}
