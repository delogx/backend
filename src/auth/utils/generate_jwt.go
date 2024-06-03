package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(claims jwt.MapClaims) (string, error) {
	var secretKey = []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
