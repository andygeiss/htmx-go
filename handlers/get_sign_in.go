package handlers

import (
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"embed"
	"log"
	"net/http"
)

func GetSignIn(efs embed.FS) http.HandlerFunc {
	te := templates.NewExecutor(efs, "assets").Parse("sign_in.html")
	return middleware.Default(func(w http.ResponseWriter, r *http.Request) {
		te.Execute(w, nil)
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
