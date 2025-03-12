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
	SearchPosts(query string) (*[]models.Post, error)
}

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{db: db}
}

// CreatePost implements PostRepository.
// creates a new post in database.
// takes a pointer as input and return error if creation fails.
func (p *PostRepositoryImpl) CreatePost(post *models.Post) error {

	if err := p.db.Create(post).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// DeletePost implements PostRepository.
// deletes post by its ID.
// takes the ID as input and return success message and error message if the post is not found or deletion fails.
func (p *PostRepositoryImpl) DeletePost(id uint) (string, error) {
	var post *models.Post
	if err := p.db.Where("id = ?", id).Delete(&post).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return "", fmt.Errorf("post with this id %v not found", id)
		}
		return "", err
	}

	return "Post deleted", nil
}

// GetAllPosts implements PostRepository.
// retrieves all posts from the database.
// returns a slice of posts. if not posts returns nil.
func (p *PostRepositoryImpl) GetAllPosts() []*models.Post {
	var posts []*models.Post

	if err := p.db.Find(&posts).Error; err != nil {
		return nil
	}

	return posts
}

// SinglePost implements PostRepository.
// retrieves single post by ID.
// take ID as input and returns posts, if post is not found it returns error.
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

// SearchPosts implements PostRepository.
// Performs a search for posts that match the query string in the title or content.
// take a query as input and returns pointer of slice of posts.
func (p *PostRepositoryImpl) SearchPosts(query string) (*[]models.Post, error) {
	var posts []models.Post

	query = "%" + query + "%"

	if err := p.db.Where("title LIKE ? OR content LIKE ?", query, query).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}
