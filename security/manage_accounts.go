package security

import (
	"errors"
	"sync"
)

type AccountManager struct {
	accounts map[string]string
	mutex    sync.Mutex
}

func (a *AccountManager) IsUsernamePasswordValid(username, password string) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if wanted, exists := a.accounts[username]; exists {
		return password == wanted
	}
	return false
}

func (a *AccountManager) RegisterAccount(username, password string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if username == "" || password == "" {
		return errors.New("The username or password cannot be empty")
	}
	if _, exists := a.accounts[username]; exists {
		return errors.New("This username is already registered")
	}
	a.accounts[username] = password
	return nil
}

func NewAccountManager() *AccountManager {
	return &AccountManager{
		accounts: map[string]string{
			"foo": "bar",
		},
	}
}

var DefaultAccountManager = NewAccountManager()
