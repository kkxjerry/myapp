package service

import (
	"encoding/csv"
	"fmt"
	"os"
)

// 导出用户为 CSV
func (service *UserService) ExportUsersCSV(filename string) error {
	users, err := service.userDao.GetAllUsers()
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入 CSV 头部
	writer.Write([]string{"ID", "Username"})

	// 写入用户数据
	for _, user := range users {
		writer.Write([]string{fmt.Sprintf("%v", user.ID), user.Username})
	}

	return nil
}
