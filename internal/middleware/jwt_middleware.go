package middleware

import (
	"Accounting/internal/auth"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const TerminalIDKey contextKey = "terminal_id"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), TerminalIDKey, claims.TerminalID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetTerminalID(r *http.Request) int64 {
	if val, ok := r.Context().Value(TerminalIDKey).(int64); ok {
		return val
	}
	return 0
}
