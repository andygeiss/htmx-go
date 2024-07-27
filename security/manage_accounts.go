package security

import "sync"

type AccountManager struct {
	accounts map[string]string
	mutex    sync.Mutex
}

func (a *AccountManager) IsUsernamePasswordValid(username, password string) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if wanted, ok := a.accounts[username]; ok {
		return password == wanted
	}
	return false
}

func NewAccountManager() *AccountManager {
	return &AccountManager{
		accounts: map[string]string{
			"foo": "bar",
		},
	}
}

var DefaultAccountManager = NewAccountManager()
