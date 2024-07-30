package authentication

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"sync"
)

type TokenManager struct {
	salt  []byte
	mutex sync.Mutex
}

func (a *TokenManager) Generate(username string) string {
	prefix := base64.RawStdEncoding.EncodeToString([]byte(username))
	suffix := base64.RawStdEncoding.EncodeToString(a.hash(username))
	return fmt.Sprintf("%s.%s", prefix, suffix)
}

func (a *TokenManager) hash(username string) []byte {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	hash := sha256.New()
	hash.Write([]byte(string(a.salt) + username))
	return hash.Sum(nil)
}

func (a *TokenManager) IsValid(token string) bool {
	parts := strings.Split(token, ".")
	username, _ := base64.RawStdEncoding.DecodeString(parts[0])
	return token == a.Generate(string(username))
}

func NewTokenManager() *TokenManager {
	salt := make([]byte, 16)
	_, _ = rand.Read(salt)
	return &TokenManager{
		salt: salt,
	}
}

var DefaultTokenManager = NewTokenManager()
