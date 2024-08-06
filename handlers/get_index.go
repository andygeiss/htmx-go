package handlers

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"log"
	"net/http"
)

type getIndexResponse struct {
	ErrorMessage string
	Token        string
}

func GetIndex(cfg *integration.Config) http.HandlerFunc {
	te := templates.NewExecutor(cfg.Efs, "assets").Parse("index.html")
	return middleware.Default(cfg, func(w http.ResponseWriter, r *http.Request) {
		te.Execute(w, getIndexResponse{ErrorMessage: "", Token: ""})
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
