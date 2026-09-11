package database

import (
	"log"

	"github.com/spf13/viper"
)

var config = viper.New()

func LoadConfig() {
	config.SetConfigFile("config.yaml")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败")
		return
	}
	log.Fatal("读取配置文件成功")
}
