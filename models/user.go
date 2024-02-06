package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string // 注意：实际应用中应该存储密码散列值
}
