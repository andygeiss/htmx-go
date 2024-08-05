package accounting

const (
	ErrorAlreadyRegistered = "Email is already registered"
	ErrorCannotBeEmpty     = "Email or password cannot be empty"
	ErrorNotRegistered     = "Email is not registered"
	ErrorWrite             = "Error during write operation"
)

type Manager interface {
	ChangePassword(email, password string) error
	IsEmailPasswordValid(email, password string) bool
	RegisterAccount(email, password string) error
}
