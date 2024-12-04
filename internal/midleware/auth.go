package middleware

import (
	"context"
	"net/http"
	"strings"

	"goresizer.com/m/internal/user"
	"goresizer.com/m/internal/utils"
)

// AuthMiddleware забезпечує авторизацію через токен
func AuthMiddleware(storage user.Storage) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := tokenParts[1]
			claims, err := utils.ParseAccessToken(tokenString)
			if err != nil {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			userID, ok := claims["user_id"].(string)
			if !ok {
				http.Error(w, "Invalid token payload", http.StatusUnauthorized)
				return
			}
			//
			//					ПЕРЕВІРКА НА ЕРРОР РОБИТЬ ЕРРОР. ЕРРОРУ БІЛЬШЕ НЕМА.
			//
			// _, err = storage.FindOne(r.Context(), userID)
			// if err != nil {
			// 	http.Error(w, "User not found", http.StatusUnauthorized)
			// 	return
			// }

			// Додаємо user_id у контекст запиту
			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
