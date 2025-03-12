package repository

import (
	"fmt"

	"github.com/ReynoldArun09/blog-application-golang/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment models.Comment) error
	GetAllComments(postID uint) (*[]models.Comment, error)
	DeleteComment(commentId uint) (string, error)
}

type CommentRepositoryImpl struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{db: db}
}

// CreateComment implements CommentRepository.
func (c *CommentRepositoryImpl) CreateComment(comment models.Comment) error {

	if err := c.db.Create(comment).Error; err != nil {
		return fmt.Errorf("failed to create comment: %w", err)
	}
	return nil
}

// DeleteComment implements CommentRepository.
func (c *CommentRepositoryImpl) DeleteComment(commentId uint) (string, error) {
	var comment *models.Comment
	if err := c.db.Where("id = ?", commentId).Delete(&comment).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return "", fmt.Errorf("comment with this id %v not found", commentId)
		}
		return "", err
	}

	return "comment deleted", nil
}

// GetAllComments implements CommentRepository.
func (c *CommentRepositoryImpl) GetAllComments(postID uint) (*[]models.Comment, error) {
	var comments []models.Comment

	if err := c.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return &comments, nil
}
