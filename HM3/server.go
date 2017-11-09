package main

import (
	"html/template"
	"net/http"

	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()
	m.Any("/", func() string {
		return "anyy"
	})

	m.Get("/", func() string {
		return "hello"
	})
	m.Get("/123.txt", func() string {
		return "trytry"
	})

	m.Get("/regist", loginControllor)
	m.RunOnAddr(":8000")
}

func loginControllor(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/htmls/regist.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}
