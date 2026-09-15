package services

import (
	"L-F/app/models"
	"L-F/app/utils"
	"L-F/configs/config"
	"L-F/configs/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 注册时检验用户是否存在
func CheckUserExistsWhenRegister(userName string) (bool, error) {
	err := database.DB.Model(&models.User{}).Where("user_name = ?", userName).First(&models.User{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// 用户注册
func Register(userName string, password string, role string, inviteCode string) (*models.User, error) {
	var user models.User

	// 将密码转化为哈希存储
	hashPassword, err := utils.Hash(password)
	if err != nil {
		return &models.User{}, ErrHashPassword
	}

	if role == "系统管理员" || role == "失物招领管理员" {
		if inviteCode != config.Config.GetString("register.admin_secret") {
			return &models.User{}, ErrNoPermission
		}
	}
	user.UserName = userName
	user.Password = hashPassword
	user.Role = role

	if flag, err := CheckUserExistsWhenRegister(userName); !flag {
		if err != nil {
			return &models.User{}, ErrUserCheckFail
		}
		err = database.DB.Model(&models.User{}).Create(&user).Error
		if err != nil {
			return &models.User{}, ErrDatabase
		}
		// 顺利注册
		return &user, nil
	}
	return &models.User{}, ErrUserExists

}

// 登陆时检验用户是否存在
func CheckUserExistsWhenLogin(userName string) (bool, *models.User, error) {
	var user models.User
	err := database.DB.Model(&models.User{}).Where("user_name = ?", userName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, &models.User{}, nil
	} else if err != nil {
		return false, &models.User{}, err
	}
	return true, &user, nil
}

// 检验密码是否正确
func CheckPassword(password1 string, password2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2))
	return err == nil
}

// 用户登录
func Login(userName string, password string) (*models.User, error) {
	flag, user, err := CheckUserExistsWhenLogin(userName)
	if err != nil {
		return user, ErrUserCheckFail
	}
	if flag {
		loginPassword := user.Password
		if CheckPassword(loginPassword, password) {
			return user, nil
		} else {
			return &models.User{}, ErrWrongPassword
		}
	} else {
		return &models.User{}, ErrUserNotFound
	}
}
