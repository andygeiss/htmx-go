package authentication_test

import (
	"andygeiss/htmx-go/usecases/authentication"
	"testing"
)

func TestDefaultManagerGenerate(t *testing.T) {
	sut := authentication.NewDefaultManager()
	if token := sut.GenerateToken("foo"); len(token) != 48 {
		t.Errorf("token length should be 48, but is %d", len(token))
	}
}

func TestDefaultManagerIsValid(t *testing.T) {
	sut := authentication.NewDefaultManager()
	if token := sut.GenerateToken("foo"); !sut.IsValidToken(token) {
		t.Error("token should be valid")
	}
	if sut.IsValidToken("some data") {
		t.Error("token should be invalid")
	}
}
