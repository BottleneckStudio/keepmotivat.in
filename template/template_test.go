package template_test

import (
	"os"
	"testing"

	temp "github.com/BottleneckStudio/keepmotivat.in/template"
)

// func TestRenderTemplate(t *testing.T) {
// 	tmpl := temp.New("../app/views/")
//
// 	helloWorld := "Hello World"
//
// 	if err := tmpl.Render(os.Stdout, "index.html", helloWorld); err != nil {
// 		t.Errorf("Must be nil, but got: %v.", err)
// 	}
//
// 	if tmpl == nil {
// 		t.Errorf("Template must not be nil. Got: %v instead.", tmpl)
// 		return
// 	}
// }

func TestParseTemplateDir(t *testing.T) {
	type TemplateData struct {
		SITENAME string
		SITEURL  string
	}

	data := TemplateData{
		SITENAME: "Template inheritance",
		SITEURL:  "https://github.com/rbo13",
	}

	tmpl, err := temp.ParseTemplateDir("../app/views")
	if err != nil {
		t.Error(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "base.html", &data)
	if err != nil {
		t.Error(err)
	}

}
