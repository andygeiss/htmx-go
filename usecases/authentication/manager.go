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
	GenerateToken(username string) string
	IsValidToken(token string) bool
}

type DefaultManager struct {
	salt  []byte
	mutex sync.Mutex
}

func (a *DefaultManager) GenerateToken(username string) string {
	prefix := base64.RawStdEncoding.EncodeToString([]byte(username))
	suffix := base64.RawStdEncoding.EncodeToString(a.hash(username))
	return fmt.Sprintf("%s.%s", prefix, suffix)
}

func (a *DefaultManager) IsValidToken(token string) bool {
	parts := strings.Split(token, ".")
	username, _ := base64.RawStdEncoding.DecodeString(parts[0])
	return token == a.GenerateToken(string(username))
}

func (a *DefaultManager) hash(username string) []byte {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	hash := sha256.New()
	hash.Write([]byte(string(a.salt) + username))
	return hash.Sum(nil)
}

func NewDefaultManager() *DefaultManager {
	salt := make([]byte, 16)
	_, _ = rand.Read(salt)
	return &DefaultManager{
		salt: salt,
	}
}
