package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config = viper.New()

func LoadConfig() {
	Config.SetConfigFile("D:/zjh_things/L-F/config.yaml")
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败")
	}
	log.Println("配置文件读取成功")
}
