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

	commentRepo := repository.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	commentController := controllers.NewCommentController(commentService)

	// Public routes (no middleware)
	mux.HandleFunc("POST /register", userController.Register)
	mux.HandleFunc("POST /login", userController.Login)

	// Routes that require authentication (middleware applied)
	mux.Handle("POST /create-post", middlewares.TokenVerifyMiddleware(http.HandlerFunc(postController.CreatePost)))
	mux.Handle("DELETE /delete-post/{id}", middlewares.TokenVerifyMiddleware(http.HandlerFunc(postController.DeletePost)))
	mux.Handle("POST /comments/create/{post_id}", middlewares.TokenVerifyMiddleware(http.HandlerFunc(commentController.CreateComment)))

	// Other routes that may not require authentication
	mux.HandleFunc("GET /all-posts", postController.GetAllPosts)
	mux.HandleFunc("GET /single-post/{id}", postController.SinglePost)
	mux.HandleFunc("GET /posts/search", postController.SearchPosts)
	mux.HandleFunc("GET /comments/{post_id}", commentController.GetAllComment)
	mux.HandleFunc("DELETE /comments/{post_id}/{comment_id}", commentController.DeleteComment)

}
