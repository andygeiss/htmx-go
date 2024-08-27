package account_test

import (
	"andygeiss/htmx-go/handlers/api/account"
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/usecases/accounting"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	WithRegistration    = true
	WithoutRegistration = false
)

type registerResponse struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

func setup(withRegistration bool) *integration.Config {
	path := "testdata/accounts.json"
	os.WriteFile(path, []byte("{}"), 0644)
	acc := accounting.NewDefaultManager(path)
	if withRegistration {
		acc.RegisterAccount("test", "test")
	}
	cfg := &integration.Config{
		AccountingManager: acc,
		ExcludedResources: []string{"/api/v1/account"},
	}
	return cfg
}

func TestRegisterSuccess(t *testing.T) {
	// Arrange
	cfg := setup(WithoutRegistration)
	data := []byte("email=test&password=test")
	r := httptest.NewRequest(http.MethodPost, "/api/v1/account", bytes.NewReader(data))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	// Act
	account.Register(cfg)(w, r)
	// Assert
	res := w.Result()
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var response registerResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("error should be nil, but got [%s]", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status should be ok, but got [%s]", res.Status)
	}
	if response.Message != "user successfully registered" {
		t.Errorf("data should be correct, but got [%s]", response.Message)
	}
}

func TestRegisterErrorAlreadyRegistered(t *testing.T) {
	// Arrange
	cfg := setup(WithRegistration)
	data := []byte("email=test&password=test")
	r := httptest.NewRequest(http.MethodPost, "/api/v1/account", bytes.NewReader(data))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	// Act
	account.Register(cfg)(w, r)
	// Assert
	res := w.Result()
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var response registerResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("error should be nil, but got [%s]", err.Error())
	}
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("status should be ok, but got [%s]", res.Status)
	}
	if response.Message != "email is already registered" {
		t.Errorf("data should be correct, but got [%s]", response.Message)
	}
}
