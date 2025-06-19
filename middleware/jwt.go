package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("your-secret-key")

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token and validate signature
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return JwtKey, nil
		}, jwt.WithLeeway(5*time.Second))

		if err != nil || !token.Valid {
			fmt.Println("Token parse error:", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userIDFloat, ok := claims[string(UserIDKey)].(float64)
		if !ok {
			http.Error(w, "Invalid token data", http.StatusUnauthorized)
			return
		}

		username, _ := claims["username"].(string)

		ctx := context.WithValue(r.Context(), UserIDKey, uint(userIDFloat))
		ctx = context.WithValue(ctx, UsernameKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
