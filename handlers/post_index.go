package handlers

import (
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"andygeiss/htmx-go/usecases/accounting"
	"andygeiss/htmx-go/usecases/authentication"
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
		if accounting.DefaultAccountManager.IsUsernamePasswordValid(username, password) {
			token = authentication.DefaultTokenManager.Generate(username)
		}
		te.Execute(w, postIndexData{Token: token})
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
