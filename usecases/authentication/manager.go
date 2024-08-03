package authentication

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"sync"
)

type Manager interface {
	GenerateToken(email string) string
	IsValidToken(token string) bool
}

type DefaultManager struct {
	salt  []byte
	mutex sync.Mutex
}

func (a *DefaultManager) GenerateToken(email string) string {
	prefix := base64.RawStdEncoding.EncodeToString([]byte(email))
	suffix := base64.RawStdEncoding.EncodeToString(a.hash(email))
	return fmt.Sprintf("%s.%s", prefix, suffix)
}

func (a *DefaultManager) IsValidToken(token string) bool {
	parts := strings.Split(token, ".")
	email, _ := base64.RawStdEncoding.DecodeString(parts[0])
	return token == a.GenerateToken(string(email))
}

func (a *DefaultManager) hash(email string) []byte {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	hash := sha256.New()
	hash.Write([]byte(string(a.salt) + email))
	return hash.Sum(nil)
}

func NewDefaultManager() *DefaultManager {
	salt := make([]byte, 16)
	_, _ = rand.Read(salt)
	return &DefaultManager{
		salt: salt,
	}
}
