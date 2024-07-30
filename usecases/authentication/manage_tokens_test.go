package authentication_test

import (
	"andygeiss/htmx-go/usecases/authentication"
	"testing"
)

func TestGenerate(t *testing.T) {
	sut := authentication.DefaultTokenManager

	if token := sut.Generate("foo"); len(token) != 48 {
		t.Errorf("Token length should be 48, but is %d", len(token))
	}
}

func TestIsValid(t *testing.T) {
	sut := authentication.NewTokenManager()

	if token := sut.Generate("foo"); !sut.IsValid(token) {
		t.Error("Token should be valid")
	}

	if sut.IsValid("some data") {
		t.Error("Token should be invalid")
	}
}
