package accounting

import (
	"context"
	"errors"
)

var (
	ErrorAlreadyRegistered = errors.New("email is already registered")
	ErrorCannotBeEmpty     = errors.New("email or password cannot be empty")
	ErrorNotRegistered     = errors.New("email is not registered")
	ErrorWrite             = errors.New("error during write operation")
)

type Manager interface {
	ChangePassword(ctx context.Context, email, password string) error
	IsEmailPasswordValid(ctx context.Context, email, password string) bool
	RegisterAccount(ctx context.Context, email, password string) error
}
