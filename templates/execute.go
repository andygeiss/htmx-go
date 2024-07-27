package templates

import (
	"embed"
	"fmt"
	"io"
	"text/template"
)

type Executor struct {
	err    error
	efs    embed.FS
	prefix string
	tmpl   *template.Template
}

func (a *Executor) Error() error {
	return a.err
}

func (a *Executor) Execute(w io.Writer, data any) {
	// Early return if an error has previously occurred.
	if a.err != nil {
		return
	}
	// Execute the template engine to merge the data.
	if err := a.tmpl.Execute(w, data); err != nil {
		a.err = err
	}
}

func (a *Executor) Parse(name string) *Executor {
	// Early return if an error has previously occurred.
	if a.err != nil {
		return a
	}
	// Parse the template from the embedded filesystem by name.
	tmpl, err := template.ParseFS(a.efs, fmt.Sprintf("%s/%s", a.prefix, name))
	// Update the template if no error has occurred.
	if err != nil {
		a.err = err
		return a
	}
	a.tmpl = tmpl
	return a
}

func NewExecutor(efs embed.FS, prefix string) *Executor {
	return &Executor{efs: efs, prefix: prefix}
}
