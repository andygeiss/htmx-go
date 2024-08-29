package authentication

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"sync"
)

type defaultManager struct {
	salt  []byte
	mutex sync.Mutex
}

func (a *defaultManager) GenerateToken(ctx context.Context, email string) string {
	prefix := base64.RawStdEncoding.EncodeToString([]byte(email))
	suffix := base64.RawStdEncoding.EncodeToString(a.hash(email))
	return fmt.Sprintf("%s.%s", prefix, suffix)
}

func (a *defaultManager) IsValidToken(ctx context.Context, token string) bool {
	parts := strings.Split(token, ".")
	email, _ := base64.RawStdEncoding.DecodeString(parts[0])
	return token == a.GenerateToken(ctx, string(email))
}

func (a *defaultManager) hash(email string) []byte {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	hash := sha256.New()
	hash.Write([]byte(string(a.salt) + email))
	return hash.Sum(nil)
}

func NewDefaultManager() Manager {
	salt := make([]byte, 16)
	_, _ = rand.Read(salt)
	return &defaultManager{
		salt: salt,
	}
}
