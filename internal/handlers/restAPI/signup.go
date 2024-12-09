package handlers

import (
	"encoding/json"
	"net/http"

	"goresizer.com/m/internal/service"
	user "goresizer.com/m/internal/storage"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpHandler(storage user.Storage, authService service.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignUpRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}

		filter := user.FindUserByFilter{
			Email: req.Email,
		}
		_, err = storage.FindOne(r.Context(), filter)
		if err == nil {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
		passhash, err := authService.HashPassword(req.Password)
		if err != nil {
			http.Error(w, "Error while hashing password", http.StatusConflict)
		}
		newUser := user.User{
			Username:     req.Username,
			Email:        req.Email,
			PasswordHash: passhash,
		}

		userID, err := storage.Create(r.Context(), newUser)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"user_id": userID})
	}
}
