package template

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	ext = ".html"
)

// Template ...
type Template struct {
	templates map[string]*template.Template
}

// Add ...
func (t Template) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	t.templates[name] = tmpl
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	if _, ok := t.templates[name]; !ok {
		// not such view
		return fmt.Errorf("no such view. (%s)", name)
	}
	return t.templates[name].Execute(w, data)
}

// RenderHTML renders as HTML.
func (t *Template) RenderHTML(w http.ResponseWriter, name string, dataTmp interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return t.Render(w, name, dataTmp)
}

// New creates a new template
func New(templatesDir string) *Template {
	result := &Template{
		templates: make(map[string]*template.Template),
	}

	layout := template.Must(template.ParseFiles(templatesDir + "layouts/base" + ext))
	template.Must(layout.ParseGlob(templatesDir + "includes/" + "*" + ext))

	dir, err := os.Open(templatesDir + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	// boot template functions.
	templateFunc := template.FuncMap{
		"toLower": func(uc string) string {
			return strings.ToLower(uc)
		},
	}

	for _, fi := range fis {
		f, err := os.Open(templatesDir + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone()).Funcs(templateFunc)
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result.templates[fi.Name()] = tmpl
	}

	return result
}
