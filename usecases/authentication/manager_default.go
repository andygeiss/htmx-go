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

func (a *defaultManager) GenerateToken(ctx context.Context, email string) (result string) {
	resultCh := make(chan string)
	go func() {
		prefix := base64.RawStdEncoding.EncodeToString([]byte(email))
		suffix := base64.RawStdEncoding.EncodeToString(a.hash(email))
		resultCh <- fmt.Sprintf("%s.%s", prefix, suffix)
	}()
	select {
	case <-ctx.Done():
	case result = <-resultCh:
	}
	return
}

func (a *defaultManager) IsValidToken(ctx context.Context, token string) (result bool) {
	resultCh := make(chan bool)
	go func() {
		parts := strings.Split(token, ".")
		email, _ := base64.RawStdEncoding.DecodeString(parts[0])
		isValid := token == a.GenerateToken(ctx, string(email))
		resultCh <- isValid
	}()
	select {
	case <-ctx.Done():
	case result = <-resultCh:
	}
	return
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
