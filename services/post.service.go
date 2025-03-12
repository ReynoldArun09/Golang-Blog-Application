package services

import (
	"github.com/ReynoldArun09/blog-application-golang/models"
	"github.com/ReynoldArun09/blog-application-golang/repository"
)

type PostService interface {
	GetAllPosts() []*models.Post
	SinglePost(id uint) (*models.Post, error)
	DeletePost(id uint) (string, error)
	CreatePost(post *models.Post) error
	SearchPosts(query string) (*[]models.Post, error)
}

type PostServiceImpl struct {
	repo repository.PostRepository
}

func (p *PostServiceImpl) SearchPosts(query string) (*[]models.Post, error) {
	posts, err := p.repo.SearchPosts(query)
	return posts, err
}

// CreatePost implements PostService.
func (p *PostServiceImpl) CreatePost(post *models.Post) error {
	return p.repo.CreatePost(post)
}

// DeletePost implements PostService.
func (p *PostServiceImpl) DeletePost(id uint) (string, error) {
	result, err := p.repo.DeletePost(id)
	return result, err

}

// GetAllPosts implements PostService.
func (p *PostServiceImpl) GetAllPosts() []*models.Post {
	return p.repo.GetAllPosts()
}

// SinglePost implements PostService.
func (p *PostServiceImpl) SinglePost(id uint) (*models.Post, error) {
	post, err := p.repo.SinglePost(id)
	return post, err
}

func NewPostService(repo repository.PostRepository) PostService {
	return &PostServiceImpl{repo: repo}
}
