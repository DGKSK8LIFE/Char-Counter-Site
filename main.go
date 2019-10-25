package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	html  *template.Template
	index = `{{define "length"}}
	Length: {{.text}}
	{{end}}	
`
)

func init() {
	html = template.Must(template.ParseGlob("index.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text := r.FormValue("text")
		htmlTwo, err := template.New("length").Parse(index)
		if err != nil {
			log.Fatal(err)
		}
		if len(text) > 0 {
			htmlTwo.ExecuteTemplate(os.Stdout, "title", nil)
		} else {
			html.ExecuteTemplate(w, "index.html", nil)
		}
	})
	http.ListenAndServe(":8080", nil)

}
