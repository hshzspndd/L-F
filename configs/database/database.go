package database

import (
	"fmt"
	"log"

	"L-F/app/models"
	"L-F/configs/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.GetString("database.user"),
		config.Config.GetString("database.password"),
		config.Config.GetString("database.host"),
		config.Config.GetInt("database.port"),
		config.Config.GetString("database.dbname"),
	)

	fmt.Println(config.Config.GetInt("database.port"))
	fmt.Println(config.Config.GetString("database.host"))
	fmt.Println(1)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("用户信息数据表创建失败")
	}

	DB = db
}
