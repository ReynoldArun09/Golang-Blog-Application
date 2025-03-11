package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ReynoldArun09/blog-application-golang/models"
	"github.com/ReynoldArun09/blog-application-golang/services"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(postService services.PostService) *PostController {
	return &PostController{postService: postService}
}

func (c *PostController) GetAllPosts(w http.ResponseWriter, r *http.Request) {

	posts := c.postService.GetAllPosts()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}
func (c *PostController) SinglePost(w http.ResponseWriter, r *http.Request) {
	post_id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid Post Id", http.StatusNotFound)
		return
	}

	post, err := c.postService.SinglePost(uint(post_id))

	if err != nil {
		http.Error(w, "failed to fetch post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)

}
func (c *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post *models.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "unable to parse data", http.StatusBadRequest)
		return
	}

	err := c.postService.CreatePost(post)

	if err != nil {
		http.Error(w, "unable to create post", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post created successfully!"})

}
func (c *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {
	post_id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid Post Id", http.StatusNotFound)
		return
	}

	post, err := c.postService.DeletePost(uint(post_id))

	if err != nil {
		http.Error(w, "failed to delete post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
