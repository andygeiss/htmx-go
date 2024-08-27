package middleware_test

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/usecases/authentication"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockupDefaultHandler(am authentication.Manager) http.HandlerFunc {
	cfg := &integration.Config{AuthenticationManager: am, ExcludedResources: []string{"/"}}
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("secure"))
	})
}

func TestDefault(t *testing.T) {
	// Arrange
	am := authentication.NewDefaultManager()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	token := am.GenerateToken("test")
	r.Header.Set("Accept-Encoding", "gzip")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	// Act
	mockupDefaultHandler(am).ServeHTTP(w, r)
	// Assert
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error should be nil, but got [%s]", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status should be ok, but got [%s]", res.Status)
	}
	if res.Header.Get("Content-Encoding") != "gzip" {
		t.Error("content-encoding should be gzip")
	}
	if len(data) != 30 {
		t.Errorf("data length should be correct, but got [%d]", len(data))
	}
}
