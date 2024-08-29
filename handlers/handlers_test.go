package handlers_test

import (
	"andygeiss/htmx-go/integration"
	"andygeiss/htmx-go/usecases/accounting"
	"andygeiss/htmx-go/usecases/authentication"
	"context"
	"embed"
)

//go:embed testdata/**
var efs embed.FS

func setup() (*integration.Config, string) {
	ctx := context.Background()
	acc := accounting.NewDefaultManager("testdata/accounts.json")
	acc.RegisterAccount(ctx, "test", "test")
	auth := authentication.NewDefaultManager()
	cfg := &integration.Config{
		AccountingManager:     acc,
		AssetsPath:            "testdata/assets",
		AuthenticationManager: auth,
		Efs:                   efs,
	}
	token := auth.GenerateToken(ctx, "test")
	return cfg, token
}
