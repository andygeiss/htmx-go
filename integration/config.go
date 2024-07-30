package integration

import (
	"andygeiss/htmx-go/usecases/accounting"
	"andygeiss/htmx-go/usecases/authentication"
	"embed"
)

type Config struct {
	AccountingManager     accounting.Manager
	AuthenticationManager authentication.Manager
	Efs                   embed.FS
	Excluded              []string // List of unsecure resources
}
