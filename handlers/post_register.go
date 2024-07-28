package handlers

import (
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"embed"
	"log"
	"net/http"
)

func PostRegister(efs embed.FS) http.HandlerFunc {
	te := templates.NewExecutor(efs, "assets").Parse("index.html")
	return middleware.Default(func(w http.ResponseWriter, r *http.Request) {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		te.Execute(w, nil)
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
