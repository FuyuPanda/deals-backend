package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("your-secret-key")

func GenerateJWT(userID uint, username string) (string, error) {
	fmt.Println(userID)
	claims := jwt.MapClaims{
		"user_id":  userID,
		"ID":       float64(userID),
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
