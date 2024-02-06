package service

import (
	"myapp/dao"
	"myapp/models"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{userDao: userDao}
}

// 验证用户登录
func (service *UserService) Authenticate(username, password string) (*models.User, bool) {
	user, err := service.userDao.FindByUsername(username)
	if err != nil || user.Password != password { // 注意: 生产环境下应使用哈希密码比较
		return nil, false
	}
	return user, true
}
