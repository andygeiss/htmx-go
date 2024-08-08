package templates_test

import (
	"andygeiss/htmx-go/templates"
	"bytes"
	"embed"
	"testing"
)

//go:embed testdata/**
var efs embed.FS

func TestExecute(t *testing.T) {
	out := new(bytes.Buffer)
	templates.NewExecutor(efs, "testdata").Parse("hello.txt").Execute(out, struct{ Name string }{Name: "Foo"})
	result := string(out.Bytes())
	if result != "Hello Foo\n" {
		t.Errorf("Result should be correct, but got [%s]", result)
	}
}
