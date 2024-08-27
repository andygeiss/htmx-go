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

func mockupValidateTokenHandler(am authentication.Manager) http.HandlerFunc {
	cfg := &integration.Config{AuthenticationManager: am, ExcludedResources: []string{"/"}}
	return middleware.ValidateToken(cfg, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("secure"))
	})
}

func TestValidateTokenSuccess(t *testing.T) {
	// Arrange
	am := authentication.NewDefaultManager()
	r := httptest.NewRequest(http.MethodGet, "/secure", nil)
	w := httptest.NewRecorder()
	token := am.GenerateToken("test")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	// Act
	mockupValidateTokenHandler(am).ServeHTTP(w, r)
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
	if len(data) != 6 {
		t.Errorf("data length should be correct, but got [%d]", len(data))
	}
}

func TestValidateTokenError(t *testing.T) {
	// Arrange
	am := authentication.NewDefaultManager()
	r := httptest.NewRequest(http.MethodGet, "/secure", nil)
	w := httptest.NewRecorder()
	// Act
	mockupValidateTokenHandler(am).ServeHTTP(w, r)
	// Assert
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error should be nil, but got [%s]", err.Error())
	}
	if res.StatusCode != http.StatusForbidden {
		t.Errorf("status should be ok, but got [%s]", res.Status)
	}
	if len(data) != 0 {
		t.Errorf("data length should be correct, but got [%d]", len(data))
	}
}
