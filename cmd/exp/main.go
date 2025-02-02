package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta Meta
}

type Meta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Doe",
		Age:  42,
		Meta: Meta{
			Visits: 5,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
