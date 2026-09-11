package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config = viper.New()

func LoadConfig() {
	Config.SetConfigFile("config.yaml")
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败")
	}
	log.Fatal("读取配置文件成功")
}
