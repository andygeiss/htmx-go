package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (a *gzipResponseWriter) Write(b []byte) (int, error) {
	return a.Writer.Write(b)
}

func newGzipResponseWriter(gw io.Writer, w http.ResponseWriter) *gzipResponseWriter {
	return &gzipResponseWriter{Writer: gw, ResponseWriter: w}
}

func CompressResponse(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Skip gzip compression if client does not support it.
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}
		// Set header and compress the body.
		w.Header().Set("Content-Encoding", "gzip")
		gw := gzip.NewWriter(w)
		defer gw.Close()
		rw := newGzipResponseWriter(gw, w)
		next.ServeHTTP(rw, r)
	}
}
