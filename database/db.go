package database

import (
	"fmt"
	"log"

	"github.com/ReynoldArun09/blog-application-golang/models"
	"github.com/ReynoldArun09/blog-application-golang/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := utils.GetEnvVariables("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to database %v", err)
	}

	fmt.Println("Database connected successfully")

	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	if err != nil {
		log.Fatalf("Failed to migrate the schema %v", err)
	}

	fmt.Println("Table created successfully or already exists")

	return db
}
