package handlers

import (
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"embed"
	"log"
	"net/http"
)

func GetHome(efs embed.FS) http.HandlerFunc {
	te := templates.NewExecutor(efs, "assets").Parse("home.html")
	return middleware.Default(func(w http.ResponseWriter, r *http.Request) {
		te.Execute(w, nil)
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
