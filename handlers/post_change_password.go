package handlers

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"log"
	"net/http"
)

type postChangePasswordData struct {
	Error   string
	Success string
}

func PostChangePassword(cfg *integration.Config) http.HandlerFunc {
	te := templates.NewExecutor(cfg.Efs, "assets").Parse("reset.html")
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		errorMessage := ""
		successMessage := "Password successfully changed"
		if err := cfg.AccountingManager.ChangePassword(email, password); err != nil {
			errorMessage = err.Error()
			successMessage = ""
		}
		te.Execute(w, postChangePasswordData{Error: errorMessage, Success: successMessage})
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
