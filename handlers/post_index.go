package handlers

import (
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/security"
	"andygeiss/htmx-go/templates"
	"embed"
	"log"
	"net/http"
)

type postIndexData struct {
	Token string
}

func PostIndex(efs embed.FS) http.HandlerFunc {
	te := templates.NewExecutor(efs, "assets").Parse("index.html")
	return middleware.Default(func(w http.ResponseWriter, r *http.Request) {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		token := ""
		if security.DefaultAccountManager.IsUsernamePasswordValid(username, password) {
			token = security.DefaultTokenManager.Generate(username)
		}
		te.Execute(w, postIndexData{Token: token})
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
