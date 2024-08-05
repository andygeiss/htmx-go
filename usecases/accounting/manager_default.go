package accounting

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type defaultManager struct {
	accounts     map[string]string
	accountsFile string
	mutex        sync.Mutex
}

func (a *defaultManager) ChangePassword(email, password string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.readAccounts()
	if _, exists := a.accounts[email]; !exists {
		return errors.New(ErrorNotRegistered)
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	a.accounts[email] = string(hash)
	if err := a.writeAccounts(); err != nil {
		return errors.New(ErrorWrite)
	}
	return nil
}

func (a *defaultManager) IsEmailPasswordValid(email, password string) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.readAccounts()
	if hash, exists := a.accounts[email]; exists {
		if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
			return false
		}
		return true
	}
	return false
}

func (a *defaultManager) RegisterAccount(email, password string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.readAccounts()
	if email == "" || password == "" {
		return errors.New(ErrorCannotBeEmpty)
	}
	if _, exists := a.accounts[email]; exists {
		return errors.New(ErrorAlreadyRegistered)
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	a.accounts[email] = string(hash)
	if err := a.writeAccounts(); err != nil {
		return errors.New(ErrorWrite)
	}
	return nil
}

func (a *defaultManager) readAccounts() error {
	data, err := os.ReadFile(a.accountsFile)
	if err != nil {
		a.accounts = map[string]string{}
		if err := os.WriteFile(a.accountsFile, []byte("{}"), 0644); err != nil {
			return err
		}
	}
	if err := json.Unmarshal(data, &a.accounts); err != nil {
		return err
	}
	return nil
}

func (a *defaultManager) writeAccounts() error {
	data, err := json.Marshal(a.accounts)
	if err != nil {
		return err
	}
	return os.WriteFile(a.accountsFile, data, 0644)
}

func NewDefaultManager(accountFile string) Manager {
	return &defaultManager{accountsFile: accountFile}
}
