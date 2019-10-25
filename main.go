package main

import (
	"html/template"
	"net/http"
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
			htmlTwo, err := template.New("").Parse(`
				<h1>Length: 
					{{range .len(text)}}	
				</h1>
			`)
			if err != nil {
				panic(err)
			}
			htmlTwo.Execute(w, htmlTwo)

		} else {
			html.ExecuteTemplate(w, "index.html", nil)
		}

	})
	http.ListenAndServe(":8080", nil)
}
