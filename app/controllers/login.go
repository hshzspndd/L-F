package controllers

import (
	"L-F/app/middles"
	"L-F/app/services"
	"L-F/app/utils"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResponseLoginData struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	Id        int       `json:"id"`
	UserName  string    `json:"username"`
	Role      string    `json:"role"`
}

func Login(c *gin.Context) {
	var loginData LoginData
	err := c.ShouldBindJSON(&loginData)
	if err != nil {
		c.Error(middles.NewError(400, "数据获取失败"))
		c.Abort()
		return
	}

	user, err := services.Login(loginData.UserName, loginData.Password)
	if err != nil {
		var bizErr *services.ResponseErrorForm
		if errors.As(err, &bizErr) {
			c.Error(middles.NewError(bizErr.Code, bizErr.Message))
		} else {
			c.Error(middles.NewError(500, "服务器内部错误"))
		}
		c.Abort()
		return
	}

	token, err, expiredAt := utils.GenerateJwt(user.UserId, user.UserName, user.Role)
	if err != nil {
		c.Error(middles.NewError(500, "登录令牌生成失败"))
		c.Abort()
		return
	}

	middles.ResponseSuccess(c, ResponseLoginData{
		Token:     token,
		ExpiredAt: expiredAt,
		Id:        user.UserId,
		UserName:  user.UserName,
		Role:      user.Role,
	})

}
