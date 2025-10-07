package middleware

import (
	"context"
	"net/http"
	"notes-app/config"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const userCtxKey contextKey = "user_id"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))

		ctx := context.WithValue(r.Context(), userCtxKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(ctx context.Context) int {
	if ctx.Value(userCtxKey) != nil {
		return ctx.Value(userCtxKey).(int)
	}
	return 0
}
