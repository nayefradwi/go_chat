package auth

import (
	"gochat/config"
	"os"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
	})
	return token.SignedString([]byte(os.Getenv(config.SECRET)))
}

// todo: missing method - decode
