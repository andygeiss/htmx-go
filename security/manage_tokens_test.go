package security_test

import (
	"andygeiss/htmx-go/security"
	. "andygeiss/htmx-go/testdata"
	"testing"
)

func TestGenerate(t *testing.T) {
	sut := security.NewTokenManager()
	token := sut.Generate("foo")

	Assert(t, "Token length should be 48", len(token), 48)
}

func TestIsValid(t *testing.T) {
	sut := security.NewTokenManager()
	token := sut.Generate("foo")

	Assert(t, "Token should be valid", sut.IsValid(token), true)
	Assert(t, "Token should be invalid", sut.IsValid("some data"), false)
}
