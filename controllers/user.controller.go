package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ReynoldArun09/blog-application-golang/models"
	"github.com/ReynoldArun09/blog-application-golang/services"
	"github.com/ReynoldArun09/blog-application-golang/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "something went wrong!!", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)

	if err := c.userService.CreateUser(user); err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	userResponse := struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	existingUser, err := c.userService.GetUser(user.Email)

	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJwt(existingUser)

	if err != nil {
		http.Error(w, "invalid credentials", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login Success",
		"token":   token,
	})
}
