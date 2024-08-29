package authentication_test

import (
	"andygeiss/htmx-go/usecases/authentication"
	"context"
	"testing"
)

func TestDefaultManagerGenerate(t *testing.T) {
	ctx := context.Background()
	sut := authentication.NewDefaultManager()
	if token := sut.GenerateToken(ctx, "foo"); len(token) != 48 {
		t.Errorf("token length should be 48, but is %d", len(token))
	}
}

func TestDefaultManagerIsValid(t *testing.T) {
	ctx := context.Background()
	sut := authentication.NewDefaultManager()
	if token := sut.GenerateToken(ctx, "foo"); !sut.IsValidToken(ctx, token) {
		t.Error("token should be valid")
	}
	if sut.IsValidToken(ctx, "some data") {
		t.Error("token should be invalid")
	}
}
