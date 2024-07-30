package main

import (
	"andygeiss/htmx-go/handlers"
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/usecases/accounting"
	"andygeiss/htmx-go/usecases/authentication"
	"embed"
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
)

//go:embed assets/**
var efs embed.FS

func main() {
	cfg := integration.Config{
		AccountingManager:     accounting.NewDefaultManager("/data/accounts.json"),
		AuthenticationManager: authentication.NewDefaultManager(),
		Efs:                   efs,
		/* The following resources does not require authentication. */
		Excluded: []string{"/", "/register", "/sign_in"},
	}
	mux := http.NewServeMux()
	mux.Handle("GET /assets/", http.FileServerFS(&cfg.Efs))
	mux.HandleFunc("GET /home", handlers.GetHome(&cfg))
	mux.HandleFunc("GET /register", handlers.GetRegister(&cfg))
	mux.HandleFunc("POST /register", handlers.PostRegister(&cfg))
	mux.HandleFunc("GET /sign_in", handlers.GetSignIn(&cfg))
	mux.HandleFunc("GET /", handlers.GetIndex(&cfg))
	mux.HandleFunc("POST /", handlers.PostIndex(&cfg))
	/* Add profiling to use Profile-guided optimization */
	mux.HandleFunc("GET /debug/pprof/", pprof.Index)
	mux.HandleFunc("GET /debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("GET /debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("GET /debug/pprof/symbol", pprof.Symbol)
	log.Printf("Start listening ...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}
