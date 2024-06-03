package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJWT(value string) (*jwt.MapClaims, bool) {
	secretKey := []byte(os.Getenv("JWT_KEY"))
	token, err := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, true
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	return &claims, ok
}
