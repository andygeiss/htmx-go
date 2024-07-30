package main

import (
	"andygeiss/htmx-go/handlers"
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/usecases/accounting"
	"andygeiss/htmx-go/usecases/authentication"
	"embed"
	"log"
	"net/http"
)

//go:embed assets/**
var efs embed.FS

func main() {
	cfg := integration.Config{
		Efs:                   efs,
		AccountingManager:     accounting.NewDefaultManager("/data/accounts.json"),
		AuthenticationManager: authentication.NewDefaultManager(),
	}
	mux := http.NewServeMux()
	mux.Handle("GET /assets/", http.FileServerFS(&cfg.Efs))
	mux.HandleFunc("GET /home", handlers.GetHome(&cfg))
	mux.HandleFunc("GET /register", handlers.GetRegister(&cfg))
	mux.HandleFunc("POST /register", handlers.PostRegister(&cfg))
	mux.HandleFunc("GET /sign_in", handlers.GetSignIn(&cfg))
	mux.HandleFunc("GET /", handlers.GetIndex(&cfg))
	mux.HandleFunc("POST /", handlers.PostIndex(&cfg))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}
