package middleware

import (
	"andygeiss/htmx-go/usecases/authentication"
	"net/http"
	"strings"
)

func ValidateToken(excluded []string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Return early if request URI is excluded.
		for _, uri := range excluded {
			if r.RequestURI == uri {
				next.ServeHTTP(w, r)
				return
			}
		}
		// Not excluded? Continue
		authHeader := r.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		token := parts[1]
		if !authentication.DefaultTokenManager.IsValid(token) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}
