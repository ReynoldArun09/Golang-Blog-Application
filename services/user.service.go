package services

import (
	"github.com/ReynoldArun09/blog-application-golang/models"
	"github.com/ReynoldArun09/blog-application-golang/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(email string) (*models.User, error)
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

// CreateUser implements UserService.
func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return u.repo.CreateUser(user)
}

// GetUser implements UserService.
func (u *UserServiceImpl) GetUser(email string) (*models.User, error) {
	user, err := u.repo.GetUser(email)
	return user, err
}
