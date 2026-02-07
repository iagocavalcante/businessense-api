package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

// UserIDKey is the context key for storing the user ID
const UserIDKey contextKey = "userID"

// JWTAuth creates a middleware that validates JWT tokens
func JWTAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				respondUnauthorized(w, "Missing auth token")
				return
			}

			parts := strings.Split(header, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				respondUnauthorized(w, "Invalid/Malformed auth token")
				return
			}

			token, err := jwt.Parse(parts[1], func(t *jwt.Token) (any, error) {
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				respondUnauthorized(w, "Malformed authentication token")
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				respondUnauthorized(w, "Token is not valid")
				return
			}

			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				respondUnauthorized(w, "Token is not valid")
				return
			}

			userID := uint(userIDFloat)
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserID extracts the user ID from the request context
func GetUserID(r *http.Request) uint {
	if id, ok := r.Context().Value(UserIDKey).(uint); ok {
		return id
	}
	return 0
}

func respondUnauthorized(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  false,
		"message": message,
	})
}
