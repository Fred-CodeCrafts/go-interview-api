package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("replace-with-your-secret")

type contextKey string

const userCtxKey = contextKey("userID")

func GenerateJWT(userID uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing auth token", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "invalid auth token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "invalid auth token claims", http.StatusUnauthorized)
			return
		}
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "invalid user id", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userCtxKey, uint(userIDFloat))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserID extracts user ID from context
func GetUserID(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value(userCtxKey).(uint)
	return userID, ok
}
