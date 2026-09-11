package services

import (
	"L-F/app/models"
	"L-F/app/utiles"
	"L-F/configs/database"

	"golang.org/x/crypto/bcrypt"
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

// 登陆时检验用户是否存在
func CheckUserExistWhenLogin(userName string) (bool, models.User) {
	var user models.User
	err := database.DB.Model(&models.User{}).Where("user_name = ?", userName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, models.User{}
	}
	return true, user
}

// 检验密码是否正确
func CheckPassword(password1 string, password2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2))
	return err == nil
}

// 用户登录
func Login(userName string, password string) (string, int, int, string) {
	flag, user := CheckUserExistWhenLogin(userName)
	if flag {
		loginPassword := user.Password
		if CheckPassword(loginPassword, password) {
			return "", 200, user.UserId, user.Role
		} else {
			return "密码错误", 403, 0, ""
		}
	} else {
		return "用户不存在", 404, 0, ""
	}
}
