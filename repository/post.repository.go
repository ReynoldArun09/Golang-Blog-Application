package repository

import (
	"fmt"

	"github.com/ReynoldArun09/blog-application-golang/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPosts() []*models.Post
	SinglePost(id uint) (*models.Post, error)
	DeletePost(id uint) (string, error)
	CreatePost(post *models.Post) error
}

type PostRepositoryImpl struct {
	db *gorm.DB
}

// CreatePost implements PostRepository.
func (p *PostRepositoryImpl) CreatePost(post *models.Post) error {

	if err := p.db.Create(post).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// DeletePost implements PostRepository.
func (p *PostRepositoryImpl) DeletePost(id uint) (string, error) {
	var post *models.Post
	if err := p.db.Where("id == ?", id).Delete(&post).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return "", fmt.Errorf("post with this id %v not found", id)
		}
		return "", err
	}

	return "Post deleted", nil
}

// GetAllPosts implements PostRepository.
func (p *PostRepositoryImpl) GetAllPosts() []*models.Post {
	var posts []*models.Post

	if err := p.db.Find(&posts).Error; err != nil {
		return nil
	}

	return posts
}

// SinglePost implements PostRepository.
func (p *PostRepositoryImpl) SinglePost(id uint) (*models.Post, error) {
	var post *models.Post
	if err := p.db.Where("id = ?", id).Find(&post).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, fmt.Errorf("post with this id %v not found", id)
		}
		return nil, err
	}

	if post == nil || post.ID == 0 {
		return nil, fmt.Errorf("post with id %v not found", id)
	}
	return post, nil
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{db: db}
}
