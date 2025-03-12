package repository

import (
	"fmt"
	"log"

	"github.com/ReynoldArun09/blog-application-golang/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(email string) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// CreateUser implements PostRepository.
// creates a new user in database.
// takes a pointer as input and return error if creation fails.
func (u *UserRepositoryImpl) CreateUser(user *models.User) error {
	if err := u.db.Create(user).Error; err != nil {
		if gorm.ErrDuplicatedKey == err {
			return fmt.Errorf("user with email %s already exists", user.Email)
		}
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetUser implements PostRepository.
// getuser receives email as input and returns user. it returns error if user not found.
func (u *UserRepositoryImpl) GetUser(email string) (*models.User, error) {
	var user models.User

	err := u.db.Where("email = ?", email).First(&user).Error

	if err != nil {

		log.Printf("Error fetching user by email: %v", err)

		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with email %s not found", email)
		}

		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return &user, nil
}
