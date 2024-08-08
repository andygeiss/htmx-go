package handlers_test

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/usecases/accounting"
	"andygeiss/htmx-go/usecases/authentication"
	"embed"
)

//go:embed testdata/**
var efs embed.FS

func setup() (*integration.Config, string) {
	acc := accounting.NewDefaultManager("testdata/accounts.json")
	acc.RegisterAccount("test", "test")
	auth := authentication.NewDefaultManager()
	cfg := &integration.Config{
		AccountingManager:     acc,
		AssetsPath:            "testdata/assets",
		AuthenticationManager: auth,
		Efs:                   efs,
	}
	token := auth.GenerateToken("test")
	return cfg, token
}
