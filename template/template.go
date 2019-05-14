package template

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
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

// New creates a new template
func New(templatesDir string) *Template {
	ins := Template{
		templates: map[string]*template.Template{},
	}

	layout := templatesDir + "layouts/base" + ext
	// adminLayout := templatesDir + "layouts/admin" + ext
	// ampLayout := templatesDir + "layouts/amp" + ext

	_, err := os.Stat(layout)
	if err != nil {
		log.Panicf("cannot find %s", layout)
		os.Exit(1)
	}

	// _, err = os.Stat(adminLayout)
	// if err != nil {
	// 	log.Printf("cannot find %s", adminLayout)
	// 	os.Exit(1)
	// }
	// _, err = os.Stat(ampLayout)
	// if err != nil {
	// 	log.Printf("cannot find %s", ampLayout)
	// 	os.Exit(1)
	// }

	partials, err := filepath.Glob(templatesDir + "partials/" + "*" + ext)
	if err != nil {
		log.Print("cannot find " + templatesDir + "partials/" + "*" + ext)
		os.Exit(1)
	}

	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}

	views, _ := filepath.Glob(templatesDir + "**/*" + ext)
	for _, view := range views {
		dir, file := filepath.Split(view)
		dir = strings.Replace(dir, templatesDir, "", 1)
		file = strings.TrimSuffix(file, ext)
		renderName := dir + file

		tmplfiles := append([]string{layout, view}, partials...)
		tmpl := template.Must(template.New(filepath.Base(layout)).Funcs(funcMap).ParseFiles(tmplfiles...))
		ins.Add(renderName, tmpl)
		log.Printf("renderName: %s, layout: %s", renderName, layout)
	}
	return &ins
}
