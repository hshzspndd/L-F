package controllers

import (
	"L-F/app/middles"
	"L-F/app/services"
	"L-F/app/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

type InformationData struct {
	Description string `json:"description" binding:"required"`
	//Images      []string `json:"images" binding:"required"`
}

type AuthorData struct {
	UserId   int    `json:"id"`
	UserName string `json:"username"`
	Role     string `json:"role"`
}
type PostData struct {
	PostType     string          `json:"type" binding:"required,oneof=lost found"`
	Title        string          `json:"title" binding:"required"`
	ContactPhone string          `json:"contact_phone" binding:"required"`
	Information  InformationData `json:"information" binding:"required"`
}

type ResponsePostData struct {
	PostId       int             `json:"post_id"`
	PostType     string          `json:"type"`
	Title        string          `json:"title"`
	ContactPhone string          `json:"contact_phone"`
	Information  InformationData `json:"information"`
	Status       string          `json:"status"`
	Author       AuthorData      `json:"author"`
}

func Post(c *gin.Context) {
	var postData PostData

	err := c.ShouldBindJSON(&postData)
	if err != nil {
		c.Error(middles.NewError(400, "数据获取失败"))
		c.Abort()
		return
	}

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
	userName := claims.UserName
	role := claims.Role
	post, err := services.CreatePost(postData.PostType, userId, postData.Title, postData.ContactPhone, postData.Information.Description)
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

	middles.ResponseSuccess(c, ResponsePostData{
		PostId:       post.PostId,
		PostType:     post.PostType,
		Title:        post.Title,
		ContactPhone: post.ContactPhone,
		Information: InformationData{
			Description: postData.Information.Description,
		},
		Status: post.Status,
		Author: AuthorData{
			UserId:   userId,
			UserName: userName,
			Role:     role,
		},
	})
}
