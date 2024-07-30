package security_test

import (
	"andygeiss/htmx-go/security"
	. "andygeiss/htmx-go/testdata"
	"os"
	"testing"
)

func TestRegisterAccount(t *testing.T) {
	path := "../testdata/test_register_account.json"
	os.WriteFile(path, []byte("{}"), 0644)
	sut := security.NewAccountManager(path)
	err := sut.RegisterAccount("foo", "bar")
	err2 := sut.RegisterAccount("foo", "bar2")

	Assert(t, "Error should be nil", err, nil)
	Assert(t, "Error should not be nil", err2 != nil, true)
	Assert(t, "Error message should be correct", err2.Error(), security.ErrorAlreadyRegistered)
}

func TestIsUsernamePasswordValid(t *testing.T) {
	path := "../testdata/test_is_username_password_valid.json"
	os.WriteFile(path, []byte("{}"), 0644)
	sut := security.NewAccountManager(path)
	_ = sut.RegisterAccount("foo", "bar")

	Assert(t, "Password should be valid", sut.IsUsernamePasswordValid("foo", "bar"), true)
	Assert(t, "Password should be invalid", sut.IsUsernamePasswordValid("foo", "bar2"), false)
}
