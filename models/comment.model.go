package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
	Post    Post   `gorm:"foreignKey:PostID" json:"post"`
}
