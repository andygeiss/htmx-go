package middleware

import (
	"andygeiss/htmx-go/integration"
	"net/http"
)

func Default(cfg *integration.Config, next http.HandlerFunc) http.HandlerFunc {
	next = LogRequest(next)
	next = ValidateToken(cfg, next)
	next = CompressResponse(next)
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}
