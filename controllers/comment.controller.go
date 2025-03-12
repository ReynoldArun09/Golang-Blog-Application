package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ReynoldArun09/blog-application-golang/middlewares"
	"github.com/ReynoldArun09/blog-application-golang/models"
	"github.com/ReynoldArun09/blog-application-golang/services"
)

type CommentController struct {
	commentService services.CommentService
	postService    services.PostService
}

func NewCommentController(commentService services.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

func (c *CommentController) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment

	userID, ok := r.Context().Value(middlewares.UserIDKey).(uint)

	if !ok || userID == 0 {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	post_id, err := strconv.Atoi(r.PathValue("post_id"))

	if err != nil {
		http.Error(w, "comment id is invalid", http.StatusNotFound)
		return
	}

	_, err = c.postService.SinglePost(uint(post_id))

	if err != nil {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	comment.UserID = userID
	comment.PostID = uint(post_id)

	err = c.commentService.CreateComment(comment)

	if err != nil {
		http.Error(w, "failed to comment", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Comment added"})

}

func (c *CommentController) DeleteComment(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(middlewares.UserIDKey).(uint)

	if !ok || userID == 0 {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	comment_id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "comment id is invalid", http.StatusNotFound)
		return
	}

	post_id, err := strconv.Atoi(r.PathValue("post_id"))

	if err != nil {
		http.Error(w, "comment id is invalid", http.StatusNotFound)
		return
	}

	_, err = c.postService.SinglePost(uint(post_id))

	if err != nil {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	}

	result, err := c.commentService.DeleteComment(uint(comment_id))

	if err != nil {
		http.Error(w, "failed to delete comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": result})

}

func (c *CommentController) GetAllComment(w http.ResponseWriter, r *http.Request) {

	post_id, err := strconv.Atoi(r.PathValue("post_id"))

	if err != nil {
		http.Error(w, "Invalid Post Id", http.StatusNotFound)
		return
	}

	comments, err := c.commentService.GetAllComments(uint(post_id))

	if err != nil {
		http.Error(w, "failed to fetch comments", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}
