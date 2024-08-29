package authentication

import "context"

type Manager interface {
	GenerateToken(ctx context.Context, email string) string
	IsValidToken(ctx context.Context, token string) bool
}
