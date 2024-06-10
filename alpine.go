package main

import (
	"html/template"
	"net/http"
)

func AlpineHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/alpine.html"))
	tmpl.Execute(w, nil)
}
