package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ReynoldArun09/blog-application-golang/middlewares"
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

	userID, ok := r.Context().Value(middlewares.UserIDKey).(uint)

	if !ok || userID == 0 {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	var post *models.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "unable to parse data", http.StatusBadRequest)
		return
	}

	post.UserID = userID

	err := c.postService.CreatePost(post)

	if err != nil {
		http.Error(w, "unable to create post", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post created successfully!"})

}
func (c *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middlewares.UserIDKey).(uint)

	if !ok || userID == 0 {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	post_id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid Post Id", http.StatusNotFound)
		return
	}

	existingPost, err := c.postService.SinglePost(uint(post_id))

	if err != nil {
		http.Error(w, "post not found", http.StatusInternalServerError)
		return
	}

	if existingPost.User.ID != userID {
		http.Error(w, "You cannot delete someone else post", http.StatusNotFound)
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

func (c *PostController) SearchPosts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	fmt.Println(query)

	if query == "" {
		http.Error(w, "Query parameter is missing", http.StatusBadRequest)
		return
	}

	posts, err := c.postService.SearchPosts(query)

	if err != nil {
		http.Error(w, "Failed to search posts", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}
