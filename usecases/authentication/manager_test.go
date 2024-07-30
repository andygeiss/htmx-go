package authentication_test

import (
	"andygeiss/htmx-go/usecases/authentication"
	"testing"
)

func TestGenerate(t *testing.T) {
	sut := authentication.NewDefaultManager()

	if token := sut.GenerateToken("foo"); len(token) != 48 {
		t.Errorf("Token length should be 48, but is %d", len(token))
	}
}

func TestIsValid(t *testing.T) {
	sut := authentication.NewDefaultManager()

	if token := sut.GenerateToken("foo"); !sut.IsValidToken(token) {
		t.Error("Token should be valid")
	}

	if sut.IsValidToken("some data") {
		t.Error("Token should be invalid")
	}
}
