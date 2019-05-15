package template

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// const (
// 	ext = ".html"
// )

// Template wraps the html/template package
type Template struct {
	*template.Template
}

// New returns a pointer custom type template.
func New(dir string) *Template {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil
	}

	tmpl, err := template.ParseFiles(paths...)

	if err != nil {
		return nil
	}

	return &Template{tmpl}
}

// Render a data by a given w.
func (tmpl *Template) Render(w io.Writer, htmlFile string, data interface{}) error {
	return tmpl.ExecuteTemplate(w, htmlFile, data)
}
