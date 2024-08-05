package authentication

type Manager interface {
	GenerateToken(email string) string
	IsValidToken(token string) bool
}
