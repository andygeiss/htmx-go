package integration

import (
	"andygeiss/htmx-go/usecases/accounting"
	"andygeiss/htmx-go/usecases/authentication"
	"embed"
)

type Config struct {
	AccountingManager     accounting.Manager
	AssetsPath            string
	AuthenticationManager authentication.Manager
	Efs                   embed.FS
	ExcludedResources     []string // This resources does not require authentication.
}
