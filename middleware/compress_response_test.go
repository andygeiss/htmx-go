package middleware_test

import (
	"andygeiss/htmx-go/middleware"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockupCompressResponseHandler() http.HandlerFunc {
	return middleware.CompressResponse(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("compressed"))
	})
}

func TestCompressResponse(t *testing.T) {
	// Arrange
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.Header.Set("Accept-Encoding", "gzip")
	// Act
	mockupCompressResponseHandler().ServeHTTP(w, r)
	// Assert
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error should be nil, but got [%s]", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, but got [%s]", res.Status)
	}
	if res.Header.Get("Content-Encoding") != "gzip" {
		t.Error("Content-Encoding should be gzip")
	}
	if len(data) != 34 {
		t.Errorf("Data length should be correct, but got [%d]", len(data))
	}
}
