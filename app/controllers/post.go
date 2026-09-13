package controllers

import (
	"L-F/app/middles"
	"L-F/app/services"
	"L-F/app/utils"

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
	PostId       int             `json:"id"`
	PostType     string          `json:"type" binding:"required,oneof=lost found"`
	Title        string          `json:"title" binding:"required"`
	ContactPhone string          `json:"contact_phone" binding:"required"`
	Information  InformationData `json:"information" binding:"required"`
	Status       string          `json:"status"`
}

type ResponsePostData struct {
	Post   PostData   `json:"post"`
	Author AuthorData `json:"author"`
}

func Post(c *gin.Context) {
	var postData PostData

	err := c.ShouldBindJSON(&postData)
	if err != nil {
		c.Error(middles.GetError(400, "数据获取失败"))
		c.Abort()
		return
	}

	v, ok := c.Get("claims")
	if !ok {
		c.Error(middles.GetError(401, "未登录"))
		c.Abort()
		return
	}

	claims, ok := v.(*utils.Claims)
	if !ok {
		c.Error(middles.GetError(401, "无效的token"))
		c.Abort()
		return
	}

	userId := claims.UserId
	userName := claims.UserName
	role := claims.Role
	message, postId, code, err := services.CreatePost(postData.PostType, userId, postData.Title, postData.ContactPhone, postData.Information.Description)
	if err != nil {
		c.Error(middles.GetError(code, message))
		c.Abort()
		return
	}

	middles.ResponseSuccess(c, ResponsePostData{
		Post: PostData{
			PostId:       postId,
			PostType:     postData.PostType,
			Title:        postData.Title,
			ContactPhone: postData.ContactPhone,
			Information: InformationData{
				Description: postData.Information.Description,
			},
			Status: "待审核",
		},
		Author: AuthorData{
			UserId:   userId,
			UserName: userName,
			Role:     role,
		},
	})
}
