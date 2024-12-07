package handlers

import (
	"encoding/json"
	"net/http"

	"goresizer.com/m/internal/storage/db"
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

func LoginHandler(storage user.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}

		user, err := storage.FindByEmail(r.Context(), req.Email)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		if !utils.VerifyPassword(req.Password, user.PasswordHash) {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

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

		response := TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
