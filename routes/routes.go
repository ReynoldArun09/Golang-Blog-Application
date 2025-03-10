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

	mux.HandleFunc("POST /register", userController.Register)
	mux.HandleFunc("POST /login", userController.Login)

}
