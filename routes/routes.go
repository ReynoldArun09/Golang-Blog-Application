package routes

import (
	"net/http"

	"github.com/ReynoldArun09/blog-application-golang/controllers"
	"github.com/ReynoldArun09/blog-application-golang/repository"
	"github.com/ReynoldArun09/blog-application-golang/services"
	"gorm.io/gorm"
)

func RegisterRoutes(mux *http.ServeMux, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	postRepo := repository.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)

	mux.HandleFunc("POST /register", userController.Register)
	mux.HandleFunc("POST /login", userController.Login)

	mux.HandleFunc("GET /all-posts", postController.GetAllPosts)
	mux.HandleFunc("GET /single-post/{id}", postController.SinglePost)
	mux.HandleFunc("DELETE /delete-post/{id}", postController.DeletePost)
	mux.HandleFunc("POST /create-post", postController.CreatePost)

}
