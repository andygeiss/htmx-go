package middleware

import (
	"andygeiss/htmx-go/integration"
	"net/http"
	"strings"
)

func ValidateToken(cfg *integration.Config, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Return early if request URI is excluded.
		for _, uri := range cfg.Excluded {
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
		if !cfg.AuthenticationManager.IsValidToken(token) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}
