package global

import (
	"github.com/spf13/viper"
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
