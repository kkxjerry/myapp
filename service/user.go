package service

import (
	"myapp/dao"
	"myapp/global"
	"myapp/models"
)

type UserService struct {
	userDao *dao.UserDao
}

var userService *UserService

func init() {
	userDao := dao.NewUserDao(global.DB)
	userService = NewUserService(userDao)
}
func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{userDao: userDao}
}

// 验证用户登录
func (service *UserService) Authenticate(username, password string) (*models.User, bool) {

	user, err := userService.userDao.FindByUsername(username)
	if err != nil || user.Password != password { // 注意: 生产环境下应使用哈希密码比较
		return nil, false
	}
	return user, true
}
