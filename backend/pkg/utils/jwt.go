package utils

import (
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(user *models.User, config *config.Config) (string, error) {
	claims := &Claims{
		ID: int64(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Server.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
