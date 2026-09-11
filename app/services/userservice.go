package services

import (
	"L-F/app/models"
	"L-F/configs/database"

	"gorm.io/gorm"
)

func CheckUserExistWhenRegister(userName string) bool {
	err := database.DB.Model(&models.User{}).Where("user_name = ?", userName).First(&models.User{}).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func Register(userName string, password string, role string) (string, int) {
	var user models.User
	user.UserName = userName
	user.Password = password
	user.Role = role
	if !CheckUserExistWhenRegister(userName) {
		err := database.DB.Model(&models.User{}).Create(&user).Error
		if err != nil {
			return "注册失败", 0
		}
		return "", user.UserId
	}
	return "用户已存在", 0
}
