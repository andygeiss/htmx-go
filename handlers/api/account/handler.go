package account

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"encoding/json"
	"net/http"
)

type registerResponse struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

func Register(cfg *integration.Config) http.HandlerFunc {
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		if err := cfg.AccountingManager.RegisterAccount(ctx, email, password); err != nil {
			data := &registerResponse{Message: err.Error(), Status: http.StatusInternalServerError}
			response, _ := json.Marshal(data)
			http.Error(w, string(response), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(&registerResponse{Message: "user successfully registered", Status: http.StatusOK})
	})
}
