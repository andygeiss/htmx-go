package accounting

import "errors"

var (
	ErrorAlreadyRegistered = errors.New("Email is already registered")
	ErrorCannotBeEmpty     = errors.New("Email or password cannot be empty")
	ErrorNotRegistered     = errors.New("Email is not registered")
	ErrorWrite             = errors.New("Error during write operation")
)

type Manager interface {
	ChangePassword(email, password string) error
	IsEmailPasswordValid(email, password string) bool
	RegisterAccount(email, password string) error
}
