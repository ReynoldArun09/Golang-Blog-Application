package routes

import (
	"net/http"

	"github.com/ReynoldArun09/blog-application-golang/controllers"
	"github.com/ReynoldArun09/blog-application-golang/middlewares"
	"github.com/ReynoldArun09/blog-application-golang/repository"
	"github.com/ReynoldArun09/blog-application-golang/services"
	"gorm.io/gorm"
)

func RegisterRoutes(mux *http.ServeMux, db *gorm.DB) {
	// Initialize repositories, services, and controllers
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	postRepo := repository.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)

	// Public routes (no middleware)
	mux.HandleFunc("/register", userController.Register)
	mux.HandleFunc("/login", userController.Login)

	// Routes that require authentication (middleware applied)
	mux.Handle("/create-post", middlewares.TokenVerifyMiddleware(http.HandlerFunc(postController.CreatePost)))
	mux.Handle("/delete-post/{id}", middlewares.TokenVerifyMiddleware(http.HandlerFunc(postController.DeletePost)))

	// Other routes that may not require authentication
	mux.HandleFunc("/all-posts", postController.GetAllPosts)
	mux.HandleFunc("/single-post/{id}", postController.SinglePost)
}
