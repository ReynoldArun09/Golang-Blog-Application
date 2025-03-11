package utils

import (
	"log"
	"os"
	"time"

	"github.com/ReynoldArun09/blog-application-golang/models"
	"github.com/dgrijalva/jwt-go"
)

func GetEnvVariables(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("warning: Environment variable %s not provided", key)
		return ""
	}

	return value

}

func GenerateJwt(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret_key := GetEnvVariables("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secret_key))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
