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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("用户信息数据表创建失败")
	}

	err = db.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatalf("失物招领信息数据表创建失败")
	}

	DB = db
}
