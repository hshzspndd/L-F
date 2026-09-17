package controllers

import (
	"L-F/app/middles"
	"L-F/app/services"
	"L-F/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Manage(c *gin.Context) {
	strPage := c.Param("page")
	page, err := strconv.Atoi(strPage)
	if err != nil {
		c.Error(middles.NewError(400, "参数获取失败"))
	}

	postType := c.Query("type")
	status := c.Query("status")

	v, ok := c.Get("claims")
	if !ok {
		c.Error(middles.NewError(401, "未登录"))
		c.Abort()
		return
	}

	claims, ok := v.(*utils.Claims)
	if !ok {
		c.Error(middles.NewError(401, "无效的token"))
		c.Abort()
		return
	}

	userId := claims.UserId

	pageResult, err := services.GetMyPosts(userId, page, postType, status)
	middles.ResponseSuccess(c, pageResult)
}
