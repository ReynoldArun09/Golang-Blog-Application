package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"size:255" json:"username"`
	Email    string    `gorm:"size:255;uniqueIndex" json:"email"`
	Password string    `gorm:"size:255" json:"password"`
	Posts    []Post    `gorm:"foreignKey:UserID" json:"posts"`
	Comments []Comment `gorm:"foreignKey:UserID" json:"comments"`
}
