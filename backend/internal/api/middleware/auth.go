package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement your authentication logic here
		// For now, we'll just pass through
		next.ServeHTTP(w, r)
	}
}
