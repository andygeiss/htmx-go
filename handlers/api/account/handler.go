package account

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"encoding/json"
	"net/http"
)

type registerData struct {
	ErrorMessage string `json:"error,omitempty"`
	Success      bool   `json:"success,omitempty"`
}

func Register(cfg *integration.Config) http.HandlerFunc {
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		errorMessage := ""
		success := true
		if err := cfg.AccountingManager.RegisterAccount(email, password); err != nil {
			errorMessage = err.Error()
			success = false
		}
		json.NewEncoder(w).Encode(&registerData{ErrorMessage: errorMessage, Success: success})
	})
}
