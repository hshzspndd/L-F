package models

type Post struct {
	PostId       int    `json:"post_id" gorm:"primarykey;autoincrement"`
	PostType     string `json:"post_type"`
	UserId       int    `json:"user_id"`
	Title        string `json:"title"`
	ContactPhone string `json:"contact_phone"`
	Description  string `json:"description"`
	//Image	 []string `json:"image"`
}
