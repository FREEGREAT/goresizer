package service

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"goresizer.com/m/internal/storage"
)

type CompressorService interface {
	Compress(imgName string, resizePercent float64) error
}

type AuthService interface {
	GenerateAccessToken(userID string, email string) (string, error)
	ValidateAccessToken(token string) (jwt.MapClaims, error)
	GenerateRefreshToken(userID string) (string, error)
	ParseRefreshToken(tokenString string) (jwt.MapClaims, error)
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) bool
}

type Storage interface {
	Create(ctx context.Context, user storage.User) (string, error)
	FindOne(ctx context.Context, customFilter storage.FindUserByFilter) (storage.User, error)
	Update(ctx context.Context, user storage.User) error
	Delete(ctx context.Context, id string) error
}
