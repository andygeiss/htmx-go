package handlers_test

import (
	"andygeiss/htmx-go/handlers"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIndex(t *testing.T) {
	// Arrange
	r := httptest.NewRequest(http.MethodGet, "/index", nil)
	w := httptest.NewRecorder()
	cfg, token := setup()
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	// Act
	handlers.GetIndex(cfg)(w, r)
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
	if string(data) != "index\n" {
		t.Errorf("data should be correct, but got [%s]", string(data))
	}
}
