package accounting_test

import (
	"andygeiss/htmx-go/usecases/accounting"
	"os"
	"testing"
)

func TestRegisterAccount(t *testing.T) {
	path := "testdata/test_register_account.json"
	os.WriteFile(path, []byte("{}"), 0644)
	sut := accounting.NewAccountManager(path)

	if err := sut.RegisterAccount("foo", "bar"); err != nil {
		t.Error("Error should be nil")
	}

	if err := sut.RegisterAccount("foo", "bar2"); err == nil {
		t.Error("Error should not be nil")
		if err.Error() != accounting.ErrorAlreadyRegistered {
			t.Error("Error message should be correct")
		}
	}
}

func TestIsUsernamePasswordValid(t *testing.T) {
	path := "testdata/test_is_username_password_valid.json"
	os.WriteFile(path, []byte("{}"), 0644)
	sut := accounting.NewAccountManager(path)

	if err := sut.RegisterAccount("foo", "bar"); err != nil {
		t.Error("Error should be nil")
	}

	if valid := sut.IsUsernamePasswordValid("foo", "bar"); !valid {
		t.Error("Password should be valid")
	}

	if valid := sut.IsUsernamePasswordValid("foo", "bar2"); valid {
		t.Error("Password should be invalid")
	}
}
