package service

import "github.com/golang-jwt/jwt/v5"

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
