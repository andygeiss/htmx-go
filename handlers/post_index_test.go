package handlers_test

import (
	"andygeiss/htmx-go/handlers"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostIndex(t *testing.T) {
	// Arrange
	data := []byte("email=test&password=test")
	r := httptest.NewRequest(http.MethodPost, "/index", bytes.NewReader(data))
	w := httptest.NewRecorder()
	cfg, token := setup()
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	// Act
	handlers.PostIndex(cfg)(w, r)
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
	if string(data) != "index\n" {
		t.Errorf("Data should be correct, but got [%s]", string(data))
	}
}
