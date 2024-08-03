package handlers

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"log"
	"net/http"
)

type postResetData struct {
	Error   string
	Success string
}

func PostReset(cfg *integration.Config) http.HandlerFunc {
	te := templates.NewExecutor(cfg.Efs, "assets").Parse("reset.html")
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		// email := r.PostFormValue("email")
		errorMessage := ""
		successMessage := "Email verification send"
		te.Execute(w, postResetData{Error: errorMessage, Success: successMessage})
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
