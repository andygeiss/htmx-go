package handlers

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"log"
	"net/http"
)

func GetRegister(cfg *integration.Config) http.HandlerFunc {
	te := templates.NewExecutor(cfg.Efs, "assets").Parse("register.html")
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		te.Execute(w, nil)
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
