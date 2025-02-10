package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{"Question 1 ?", "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fugiat, minima."},
		{"Question 2 ?", "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fugiat, minima."},
		{"Question 3 ?", "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fugiat, minima."},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
