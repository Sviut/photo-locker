package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	tpl := template.New(pattern[0])
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return `<input type=hidden" />`
		},
	})

	tpl, err := tpl.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		HtmlTpl: tpl,
	}, nil
}

type Template struct {
	HtmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := t.HtmlTpl.Execute(w, data); err != nil {
		log.Printf("Error execute template: %v", err)
		http.Error(w, "Error execute template", http.StatusInternalServerError)
		return
	}
}
