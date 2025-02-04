package controllers

import (
	"net/http"

	"github.com/sviut/photo-locker/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{"Question 1 ?", "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fugiat, minima."},
		{"Question 2 ?", "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fugiat, minima."},
		{"Question 3 ?", "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fugiat, minima."},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
