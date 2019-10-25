package main

import (
	"html/template"
	"log"
	"net/http"
)

var html *template.Template

func init() {
	html = template.Must(template.ParseGlob("index.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text := r.FormValue("text")
		log.Printf("characters: %v", len(text))
		html.ExecuteTemplate(w, "index.html", nil)
	})
	http.ListenAndServe(":8080", nil)
}
