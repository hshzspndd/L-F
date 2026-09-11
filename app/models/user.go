package models

type User struct {
	UserId   int    `json:"user_id" gorm:"primarykey;autoIncrement"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"check:role IN ('系统管理员','失物招领管理员','普通用户')"`
}
