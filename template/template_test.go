package template_test

import (
	"os"
	"testing"

	temp "github.com/BottleneckStudio/keepmotivat.in/template"
)

func TestParseTemplateDir(t *testing.T) {
	type Fixture struct {
		Fullname string
		Message  string
		URL      string
	}

	data := Fixture{
		Fullname: "Richard Burk",
		Message:  "Hello",
		URL:      "https://github.com/rbo13",
	}

	tmpl := temp.New("../app/views/")

	if tmpl == nil {
		t.Fatalf("Creating instance of template should not be nil, but got an error: [ %v ] instead.", tmpl)
	}

	if err := tmpl.Render(os.Stdout, "hello/hello", &data); err != nil {
		t.Fatalf("Expects to render an html output, but got an error: [ %v ] instead.", err)
	}
}

func BenchmarkRender(b *testing.B) {
	b.ReportAllocs()
	tmpl := temp.New("../app/views/")

	for i := 0; i < b.N; i++ {
		if err := tmpl.Render(os.Stdout, "hello/hello", nil); err != nil {
			panic(err)
		}
	}
}
