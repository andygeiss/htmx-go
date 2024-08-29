package accounting_test

import (
	"andygeiss/htmx-go/usecases/accounting"
	"context"
	"os"
	"testing"
)

func TestDefaultManagerChangePassword(t *testing.T) {
	ctx := context.Background()
	path := "testdata/test_default_manager_change_password.json"
	os.WriteFile(path, []byte("{}"), 0644)
	sut := accounting.NewDefaultManager(path)
	sut.RegisterAccount(ctx, "foo", "bar")
	if err := sut.ChangePassword(ctx, "foo", "bar2"); err != nil {
		t.Error("error should be nil")
	}
}

func TestDefaultManagerIsEmailPasswordValid(t *testing.T) {
	ctx := context.Background()
	path := "testdata/test_default_manager_is_email_password_valid.json"
	os.WriteFile(path, []byte("{}"), 0644)
	sut := accounting.NewDefaultManager(path)
	if err := sut.RegisterAccount(ctx, "foo", "bar"); err != nil {
		t.Error("error should be nil")
	}
	if valid := sut.IsEmailPasswordValid(ctx, "foo", "bar"); !valid {
		t.Error("password should be valid")
	}
	if valid := sut.IsEmailPasswordValid(ctx, "foo", "bar2"); valid {
		t.Error("password should be invalid")
	}
}

func TestDefaultManagerRegisterAccount(t *testing.T) {
	ctx := context.Background()
	path := "testdata/test_default_manager_register_account.json"
	os.WriteFile(path, []byte("{}"), 0644)
	sut := accounting.NewDefaultManager(path)
	if err := sut.RegisterAccount(ctx, "foo", "bar"); err != nil {
		t.Error("error should be nil")
	}
	if err := sut.RegisterAccount(ctx, "foo", "bar2"); err == nil {
		t.Error("error should not be nil")
		if err != accounting.ErrorAlreadyRegistered {
			t.Error("error message should be correct")
		}
	}
}
