package main

import (
	"html/template"
	"net/http"
)

var (
	html    *template.Template
	htmlTwo *template.Template
)

func init() {
	html = template.Must(template.ParseGlob("index.html"))
	htmlTwo = template.Must(template.ParseGlob("length.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text := r.FormValue("text")
		if len(text) > 0 {
			htmlTwo.ExecuteTemplate(w, "length.html", nil)
		} else {
			html.ExecuteTemplate(w, "index.html", nil)
		}

	})
	http.ListenAndServe(":8080", nil)
}
