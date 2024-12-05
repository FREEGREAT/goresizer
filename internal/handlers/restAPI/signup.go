package handlers

import (
	"encoding/json"
	"net/http"

	"goresizer.com/m/internal/user"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpHandler(storage user.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignUpRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}

		_, err = storage.FindByEmail(r.Context(), req.Email)
		if err == nil {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
		newUser := user.User{
			Username:     req.Username,
			Email:        req.Email,
			PasswordHash: req.Password,
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
