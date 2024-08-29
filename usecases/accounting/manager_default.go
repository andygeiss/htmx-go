package accounting

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type defaultManager struct {
	accounts     map[string]string
	accountsFile string
	mutex        sync.Mutex
}

func (a *defaultManager) ChangePassword(ctx context.Context, email, password string) (err error) {
	doneCh := make(chan bool)
	errCh := make(chan error)
	go func() {
		a.mutex.Lock()
		defer a.mutex.Unlock()
		a.readAccounts()
		if _, exists := a.accounts[email]; !exists {
			errCh <- ErrorNotRegistered
			return
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
		a.accounts[email] = string(hash)
		if err := a.writeAccounts(); err != nil {
			errCh <- ErrorWrite
		}
		doneCh <- true
	}()
	select {
	case <-ctx.Done():
	case <-doneCh:
	case err = <-errCh:
		return
	}
	return
}

func (a *defaultManager) IsEmailPasswordValid(ctx context.Context, email, password string) (result bool) {
	resultCh := make(chan bool)
	errCh := make(chan error)
	go func() {
		a.mutex.Lock()
		defer a.mutex.Unlock()
		a.readAccounts()
		if hash, exists := a.accounts[email]; exists {
			if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
				errCh <- err
				return
			}
			resultCh <- true
			return
		}
		resultCh <- false
	}()
	select {
	case <-ctx.Done():
	case result = <-resultCh:
		return
	case <-errCh:
	}
	return
}

func (a *defaultManager) RegisterAccount(ctx context.Context, email, password string) (err error) {
	doneCh := make(chan bool)
	errCh := make(chan error)
	go func() {
		a.mutex.Lock()
		defer a.mutex.Unlock()
		a.readAccounts()
		if email == "" || password == "" {
			errCh <- ErrorCannotBeEmpty
			return
		}
		if _, exists := a.accounts[email]; exists {
			errCh <- ErrorAlreadyRegistered
			return
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
		a.accounts[email] = string(hash)
		if err := a.writeAccounts(); err != nil {
			errCh <- ErrorWrite
			return
		}
		doneCh <- true
	}()
	select {
	case <-ctx.Done():
	case <-doneCh:
	case err = <-errCh:
		return
	}
	return
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
