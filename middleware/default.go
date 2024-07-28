package middleware

import (
	"net/http"
)

var excluded = []string{
	"/", "/sign_in",
}

func Default(next http.HandlerFunc) http.HandlerFunc {
	next = LogRequest(next)
	next = ValidateToken(excluded, next)
	next = CompressResponse(next)
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}
