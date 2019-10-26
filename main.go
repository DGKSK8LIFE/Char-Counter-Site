package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var (
	html *template.Template
)

func init() {
	html = template.Must(template.ParseGlob("index.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text := r.FormValue("text")
		if len(text) > 0 {
			fmt.Fprintf(w, "Char Count: %d", len(countChars(text)))
		} else {
			html.ExecuteTemplate(w, "index.html", nil)
		}

	})
	http.ListenAndServe(":8080", nil)

}

func countChars(text string) string {
	return strings.Replace(text, " ", "", -1)
}
