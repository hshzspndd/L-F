package services

import (
	"L-F/app/models"
	"L-F/app/utiles"
	"L-F/configs/database"

	"gorm.io/gorm"
)

// 注册时检验用户是否存在
func CheckUserExistWhenRegister(userName string) bool {
	err := database.DB.Model(&models.User{}).Where("user_name = ?", userName).First(&models.User{}).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// 用户注册
func Register(userName string, password string, role string) (string, int, int) {
	var user models.User

	// 将密码转化为哈希存储
	hash_password, err := utiles.Hash(password)
	if err != nil {
		return "密码加密失败", 0, 500
	}

	user.UserName = userName
	user.Password = hash_password
	user.Role = role

	if !CheckUserExistWhenRegister(userName) {
		err = database.DB.Model(&models.User{}).Create(&user).Error
		if err != nil {
			return "注册失败", 0, 500
		}

		// 顺利注册
		return "", user.UserId, 200
	}
	return "用户已存在", 0, 409
}
