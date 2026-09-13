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
func Register(userName string, password string, role string, inviteCode string) (models.User, *ResponseErrorForm, bool) {
	var user models.User

	// 将密码转化为哈希存储
	hashPassword, err := utils.Hash(password)
	if err != nil {
		return models.User{}, ErrHashPassword, false
	}

	if role == "系统管理员" || role == "失物招领管理员" {
		if inviteCode != config.Config.GetString("register.admin_secret") {
			return models.User{}, ErrNoPermission, false
		}
	}
	user.UserName = userName
	user.Password = hashPassword
	user.Role = role

	if flag, err := CheckUserExistsWhenRegister(userName); !flag {
		if err != nil {
			return models.User{}, ErrUserCheckFail, false
		}
		err = database.DB.Model(&models.User{}).Create(&user).Error
		if err != nil {
			return models.User{}, ErrDatabase, false
		}
		// 顺利注册
		return user, &ResponseErrorForm{}, true
	}
	return models.User{}, ErrUserExists, false

}

// 登陆时检验用户是否存在
func CheckUserExistsWhenLogin(userName string) (bool, error, models.User) {
	var user models.User
	err := database.DB.Model(&models.User{}).Where("user_name = ?", userName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil, models.User{}
	} else if err != nil {
		return false, err, models.User{}
	}
	return true, nil, user
}

// 检验密码是否正确
func CheckPassword(password1 string, password2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2))
	return err == nil
}

// 用户登录
func Login(userName string, password string) (models.User, *ResponseErrorForm, bool) {
	flag, err, user := CheckUserExistsWhenLogin(userName)
	if err != nil {
		return user, ErrUserCheckFail, false
	}
	if flag {
		loginPassword := user.Password
		if CheckPassword(loginPassword, password) {
			return user, &ResponseErrorForm{}, true
		} else {
			return models.User{}, ErrWrongPassword, false
		}
	} else {
		return models.User{}, ErrUserNotFound, false
	}
}

// 创建失物招领信息
func CreatePost(postType string, userId int, title string, contactPhone string, description string) (models.Post, error) {
	var post models.Post
	post.PostType = postType
	post.UserId = userId
	post.Title = title
	post.ContactPhone = contactPhone
	post.Description = description
	post.Status = "待审核"
	err := database.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, ErrDatabase
	}

	return post, nil

}
