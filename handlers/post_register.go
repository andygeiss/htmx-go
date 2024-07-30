package handlers

import (
	"andygeiss/htmx-go/middleware"
	"andygeiss/htmx-go/templates"
	"andygeiss/htmx-go/usecases/accounting"
	"embed"
	"log"
	"net/http"
)

type postRegisterData struct {
	Error   string
	Success string
}

func PostRegister(efs embed.FS) http.HandlerFunc {
	te := templates.NewExecutor(efs, "assets").Parse("register.html")
	return middleware.Default(func(w http.ResponseWriter, r *http.Request) {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		errorMessage := ""
		successMessage := "Account successfully created"
		if err := accounting.DefaultAccountManager.RegisterAccount(username, password); err != nil {
			errorMessage = err.Error()
			successMessage = ""
		}
		te.Execute(w, postRegisterData{Error: errorMessage, Success: successMessage})
		if te.Error() != nil {
			log.Println(te.Error())
		}
	})
}
