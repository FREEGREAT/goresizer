package handlers

import (
	"encoding/json"
	"net/http"

	"goresizer.com/m/internal/user"
	"goresizer.com/m/internal/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// LoginHandler приймає `user.Storage` і повертає http.HandlerFunc.
func LoginHandler(storage user.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		// Парсинг JSON із запиту
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}

		// Знаходимо користувача за email
		user, err := storage.FindByEmail(r.Context(), req.Email)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Перевірка пароля (без хешування)
		if user.PasswordHash != req.Password {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Генерація токенів
		accessToken, err := utils.GenerateAccessToken(user.ID, user.Email)
		if err != nil {
			http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
			return
		}

		refreshToken, err := utils.GenerateRefreshToken(user.ID)
		if err != nil {
			http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
			return
		}

		// Формуємо відповідь
		response := TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		// Відправляємо токени клієнту
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
