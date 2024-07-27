package middleware

import (
	"log"
	"net/http"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(`INFO %20s "%s %s"`, r.RemoteAddr, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
