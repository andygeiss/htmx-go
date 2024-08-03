package handlers

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"log"
	"net/http"
)

type postIndexData struct {
	Token string
}

func PostIndex(cfg *integration.Config) http.HandlerFunc {
	te := templates.NewExecutor(cfg.Efs, "assets").Parse("index.html")
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		token := ""
		if cfg.AccountingManager.IsEmailPasswordValid(email, password) {
			token = cfg.AuthenticationManager.GenerateToken(email)
		}
		te.Execute(w, postIndexData{Token: token})
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
