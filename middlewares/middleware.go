package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ReynoldArun09/blog-application-golang/utils"
	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const UserIDKey contextKey = "userID"

// middleware to validate token and inject userid in ctx
func TokenVerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		if tokenString == "" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token uses the expected signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return the secret key used to sign the JWT
			secret_key := utils.GetEnvVariables("SECRET_KEY")
			return []byte(secret_key), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userIDStr := fmt.Sprintf("%v", claims["sub"])
		userID, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, uint(userID))

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
