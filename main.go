package main

import (
	"andygeiss/htmx-go/handlers"
	"andygeiss/htmx-go/handlers/api/account"
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
		AssetsPath:            "assets",
		AuthenticationManager: authentication.NewDefaultManager(),
		Efs:                   efs,
		ExcludedResources:     []string{"/", "/index.html", "/sign_in.html", "/api/v1/account"},
	}
	mux := http.NewServeMux()
	mux.Handle("GET /assets/", http.FileServerFS(&cfg.Efs))
	/* Add API resources */
	mux.HandleFunc("POST /api/v1/account", account.Register(&cfg))
	/* Add basic accounting and authentication */
	mux.HandleFunc("GET /", handlers.GetIndex(&cfg))
	mux.HandleFunc("GET /home.html", handlers.GetHome(&cfg))
	mux.HandleFunc("GET /sign_in.html", handlers.GetSignIn(&cfg))
	mux.HandleFunc("GET /index.html", handlers.GetIndex(&cfg))
	mux.HandleFunc("POST /index.html", handlers.PostIndex(&cfg))
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
