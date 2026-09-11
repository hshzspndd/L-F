package controllers

import (
	"L-F/app/middles"
	"L-F/app/services"

	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	UserName string `json:"username" binding:"required,min=3,max=15"`
	Password string `json:"password" binding:"required,min=6,max=15"`
	Role     string `json:"role" binding:"required,oneof=系统管理员 失物招领管理员 普通用户"`
}

type ResponseRegisterData struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Role     string `json:"role"`
}

func Register(c *gin.Context) {
	var register_data RegisterData
	err := c.ShouldBindJSON(&register_data)
	if err != nil {
		c.Error(middles.GetError(400, "数据获取失败"))
		return
	}

	msg, id, code := services.Register(register_data.UserName, register_data.Password, register_data.Role)
	if msg != "" {
		c.Error(middles.GetError(code, msg))
		return
	}

	middles.ResponseSuccess(c, ResponseRegisterData{
		Id:       id,
		UserName: register_data.UserName,
		Role:     register_data.Role,
	})
}
