package main

import (
	"andygeiss/htmx-go/handlers"
	"embed"
	"log"
	"net/http"
)

//go:embed assets/**
var efs embed.FS

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /assets/", http.FileServerFS(efs))
	mux.HandleFunc("GET /home", handlers.GetHome(efs))
	mux.HandleFunc("GET /register", handlers.GetRegister(efs))
	mux.HandleFunc("GET /sign_in", handlers.GetSignIn(efs))
	mux.HandleFunc("GET /", handlers.GetIndex(efs))
	mux.HandleFunc("POST /", handlers.PostIndex(efs))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}
