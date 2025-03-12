package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string    `gorm:"size:255" json:"title"`
	Content  string    `gorm:"type:text" json:"content"`
	Username string    `gorm:"size:255" json:"username"`
	UserID   uint      `json:"user_id"`
	User     User      `gorm:"foreignKey:UserID" json:"user"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments"`
}
