package views

import (
	"fmt"
	"github.com/gorilla/csrf"
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

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := t.HtmlTpl.Clone()
	if err != nil {
		log.Printf("Error cloning template: %v", err)
		http.Error(w, "Error rendering page.", http.StatusInternalServerError)
		return
	}
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
	})
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = tpl.Execute(w, data)

	if err != nil {
		log.Printf("Error execute template: %v", err)
		http.Error(w, "Error execute template", http.StatusInternalServerError)
		return
	}
}
