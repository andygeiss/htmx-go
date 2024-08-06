package handlers

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"log"
	"net/http"
)

type postIndexResponse struct {
	ErrorMessage string
	Token        string
}

func PostIndex(cfg *integration.Config) http.HandlerFunc {
	te := templates.NewExecutor(cfg.Efs, "assets").Parse("index.html")
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		errorMessage := ""
		token := ""
		if cfg.AccountingManager.IsEmailPasswordValid(email, password) {
			token = cfg.AuthenticationManager.GenerateToken(email)
		} else {
			errorMessage = "Incorrect email or password"
		}
		te.Execute(w, postIndexResponse{ErrorMessage: errorMessage, Token: token})
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
