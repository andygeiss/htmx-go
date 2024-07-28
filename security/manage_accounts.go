package security

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type AccountManager struct {
	accounts    map[string]string
	accountFile string
	mutex       sync.Mutex
}

func (a *AccountManager) IsUsernamePasswordValid(username, password string) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.readAccounts()
	if hash, exists := a.accounts[username]; exists {
		_ = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		return true
	}
	return false
}

func (a *AccountManager) RegisterAccount(username, password string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.readAccounts()
	if username == "" || password == "" {
		return errors.New("The username or password cannot be empty")
	}
	if _, exists := a.accounts[username]; exists {
		return errors.New("This username is already registered")
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	a.accounts[username] = string(hash)
	if err := a.writeAccounts(); err != nil {
		return errors.New("Error during write operation")
	}
	return nil
}

func (a *AccountManager) readAccounts() error {
	data, err := os.ReadFile(a.accountFile)
	if err != nil {
		a.accounts = map[string]string{}
		if err := os.WriteFile(a.accountFile, []byte("{}"), 0644); err != nil {
			return err
		}
	}
	if err := json.Unmarshal(data, &a.accounts); err != nil {
		return err
	}
	return nil
}

func (a *AccountManager) writeAccounts() error {
	data, err := json.Marshal(a.accounts)
	if err != nil {
		return err
	}
	return os.WriteFile(a.accountFile, data, 0644)
}

func NewAccountManager(accountFile string) *AccountManager {
	return &AccountManager{accountFile: accountFile}
}

var DefaultAccountManager = NewAccountManager("data/accounts.json")
