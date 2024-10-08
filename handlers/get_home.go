package handlers

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"log"
	"net/http"
)

func GetHome(cfg *integration.Config) http.HandlerFunc {
	te := templates.NewExecutor(cfg.Efs, cfg.AssetsPath).Parse("home.html")
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		te.Execute(w, nil)
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
