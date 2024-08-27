package accounting

import "errors"

var (
	ErrorAlreadyRegistered = errors.New("email is already registered")
	ErrorCannotBeEmpty     = errors.New("email or password cannot be empty")
	ErrorNotRegistered     = errors.New("email is not registered")
	ErrorWrite             = errors.New("error during write operation")
)

type Manager interface {
	ChangePassword(email, password string) error
	IsEmailPasswordValid(email, password string) bool
	RegisterAccount(email, password string) error
}
