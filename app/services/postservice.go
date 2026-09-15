package services

import (
	"L-F/app/models"
	"L-F/configs/database"
)

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
