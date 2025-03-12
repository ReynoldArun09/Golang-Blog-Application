package services

import (
	"github.com/ReynoldArun09/blog-application-golang/models"
	"github.com/ReynoldArun09/blog-application-golang/repository"
)

type CommentService interface {
	CreateComment(comment models.Comment) error
	GetAllComments(postID uint) (*[]models.Comment, error)
	DeleteComment(commentId uint) (string, error)
}

type CommentServiceImpl struct {
	commentRepository repository.CommentRepository
}

// CreateComment implements CommentService.
func (c *CommentServiceImpl) CreateComment(comment models.Comment) error {
	return c.commentRepository.CreateComment(comment)
}

// DeleteComment implements CommentService.
func (c *CommentServiceImpl) DeleteComment(commentId uint) (string, error) {
	return c.commentRepository.DeleteComment(commentId)
}

// GetAllComments implements CommentService.
func (c *CommentServiceImpl) GetAllComments(postID uint) (*[]models.Comment, error) {
	return c.commentRepository.GetAllComments(postID)
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &CommentServiceImpl{commentRepository: repo}
}
