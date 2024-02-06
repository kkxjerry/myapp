package global

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Configuration struct {
	Server struct {
		Port string
	}
	Database struct {
		Type     string
		User     string
		Password string
		Name     string
		Host     string
		Port     string
	}
	Security struct {
		SecretKey string
	}
}

var Config Configuration
var DB *gorm.DB

func LoadConfig() {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")      // 查找配置文件所在的路径

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

}
func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.Database.User,
		Config.Database.Password,
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.Name)

	DB, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
