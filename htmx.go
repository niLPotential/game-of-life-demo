package main

import (
	"html/template"
	"net/http"
)

func HtmxHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/htmx.html"))
	tmpl.Execute(w, nil)
}
