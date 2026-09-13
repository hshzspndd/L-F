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
func Register(userName string, password string, role string, inviteCode string) (string, int, int) {
	var user models.User

	// 将密码转化为哈希存储
	hash_password, err := utils.Hash(password)
	if err != nil {
		return "密码加密失败", 0, 500
	}

	if role == "系统管理员" || role == "失物招领管理员" {
		if inviteCode != config.Config.GetString("register.admin_secret") {
			return "没有权限", 0, 403
		}
	}
	user.UserName = userName
	user.Password = hash_password
	user.Role = role

	if flag, err := CheckUserExistsWhenRegister(userName); !flag {
		if err != nil {
			return "用户信息校验失败", 0, 500
		}
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
func Login(userName string, password string) (string, int, int, string) {
	flag, err, user := CheckUserExistsWhenLogin(userName)
	if err != nil {
		return "用户信息校验失败", 500, 0, ""
	}
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

// 创建失物招领信息
func CreatePost(postType string, userId int, title string, contactPhone string, description string) (string, int, int, error) {
	var post models.Post
	post.PostType = postType
	post.UserId = userId
	post.Title = title
	post.ContactPhone = contactPhone
	post.Description = description
	post.Status = "待审核"
	err := database.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return "数据库出错", 0, 500, err
	}

	return "发布成功", post.PostId, 200, nil

}
