package handlers

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"log"
	"net/http"
)

type postRegisterData struct {
	Error   string
	Success string
}

func PostRegister(cfg *integration.Config) http.HandlerFunc {
	te := templates.NewExecutor(cfg.Efs, "assets").Parse("register.html")
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		errorMessage := ""
		successMessage := "Account successfully created"
		if err := cfg.AccountingManager.RegisterAccount(username, password); err != nil {
			errorMessage = err.Error()
			successMessage = ""
		}
		te.Execute(w, postRegisterData{Error: errorMessage, Success: successMessage})
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
