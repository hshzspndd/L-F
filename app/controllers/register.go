package controllers

import (
	"L-F/app/middles"
	"L-F/app/services"
	"errors"

	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	UserName   string `json:"username" binding:"required,min=3,max=15"`
	Password   string `json:"password" binding:"required,min=6,max=15"`
	Role       string `json:"role" binding:"required,oneof=系统管理员 失物招领管理员 普通用户"`
	InviteCode string `json:"invite_code"`
}

type ResponseRegisterData struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Role     string `json:"role"`
}

func Register(c *gin.Context) {
	var registerData RegisterData
	err := c.ShouldBindJSON(&registerData)
	if err != nil {
		c.Error(middles.GetError(400, "数据获取失败"))
		c.Abort()
		return
	}

	user, err := services.Register(registerData.UserName, registerData.Password, registerData.Role, registerData.InviteCode)
	if err != nil {
		var bizErr *services.ResponseErrorForm
		if errors.As(err, &bizErr) {
			c.Error(middles.GetError(bizErr.Code, bizErr.Message))
		} else {
			c.Error(middles.GetError(500, "服务器内部错误"))
		}
		c.Abort()
		return
	}

	middles.ResponseSuccess(c, ResponseRegisterData{
		Id:       user.UserId,
		UserName: user.UserName,
		Role:     user.Role,
	})
}
