package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
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
			startTime := time.Now()
			fmt.Fprintf(w, "<h1 style=\"text-align: center;\">Counted %d Characters in %fs</h1>", len(countChars(text)), time.Since(startTime).Seconds())
		} else {
			html.ExecuteTemplate(w, "index.html", nil)
		}

	})
	http.ListenAndServe(":8080", nil)

}

func countChars(text string) string {
	return strings.Replace(text, " ", "", -1)
}
