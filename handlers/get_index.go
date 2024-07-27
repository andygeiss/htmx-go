package handlers

import (
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"embed"
	"log"
	"net/http"
)

type getIndexData struct {
	Token string
}

func GetIndex(efs embed.FS) http.HandlerFunc {
	te := templates.NewExecutor(efs, "assets").Parse("index.html")
	return middleware.Default(func(w http.ResponseWriter, r *http.Request) {
		te.Execute(w, getIndexData{Token: ""})
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
