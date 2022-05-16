package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/nayefradwi/go_chat/common/errorHandling"
)

func GenerateToken(userId int, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
	})
	return token.SignedString([]byte(secret))
}

func DecodeAccessToken(tokenString string, secret string) (int, *errorHandling.BaseError) {
	if isParsed, token := verifyToken(tokenString, secret); isParsed {
		claims := parseToken(token)
		if val, ok := claims["id"]; ok {
			userId := int(val.(float64))
			return userId, nil
		}
	}
	return -1, errorHandling.NewUnAuthorizedError()
}

func verifyToken(tokenString string, secret string) (bool, *jwt.Token) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errorHandling.NewUnAuthorizedError()
		}
		return []byte(secret), nil
	})
	return err == nil && token.Valid, token
}

func parseToken(token *jwt.Token) jwt.MapClaims {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		return claims
	}
	return jwt.MapClaims{}
}
