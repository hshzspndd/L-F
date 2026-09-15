// 定义返回给前端的结构体
package services

import (
	"L-F/app/models"
	"L-F/configs/database"
)

type InformationData struct {
	Description string `json:"description" binding:"required"`
	//Images      []string `json:"images" binding:"required"`
}

type ResponsePageData struct {
	PostId       int             `json:"post_id"`
	PostType     string          `json:"type"`
	Title        string          `json:"title"`
	ContactPhone string          `json:"contact_phone"`
	Information  InformationData `json:"information"`
	Status       string          `json:"status"`
}
type PageResult struct {
	List     []ResponsePageData `json:"list"`
	Total    int                `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"pagesize"`
}

// GetMyPosts 分页查询
func GetMyPosts(userId int, page int, postType string, status string) (*PageResult, error) {
	var posts []models.Post
	var total int64
	var pageSize int = 10
	// 1. 构建基础查询
	query := database.DB.Model(&models.Post{}).Where("user_id = ?", userId)

	// 如果前端传了状态过滤（例如 ?status=待审核）
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 2. 先查总条数（用于前端算总页数）
	if err := query.Count(&total).Error; err != nil {
		return nil, ErrDatabase
	}

	// 3. 计算 offset，执行分页查询
	offset := (page - 1) * pageSize
	err := query.Order("post_id DESC").Offset(offset).Limit(pageSize).Find(&posts).Error // 按时间倒序，最新的在最上面

	if err != nil {
		return nil, ErrDatabase
	}

	list := make([]ResponsePageData, 0, len(posts))
	for _, post := range posts {
		item := ResponsePageData{
			PostId:       post.PostId,
			PostType:     post.PostType,
			Title:        post.Title,
			ContactPhone: post.ContactPhone,
			Information: InformationData{
				Description: post.Description,
			},
			Status: post.Status,
		}

		list = append(list, item)
	}
	return &PageResult{
		List:     list,
		Total:    int(total),
		Page:     page,
		PageSize: len(list),
	}, nil
}
