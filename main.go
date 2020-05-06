package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var templates *template.Template

func main() {

	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	templates = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", askiiHendler)
	http.ListenAndServe(":8080", nil)
}

func askiiHendler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(404)
		fmt.Fprintf(w, "404 not found")
		return
	}

	switch {
	case r.Method == "GET":
		templates.ExecuteTemplate(w, "index.html", nil)
	case r.Method == "POST":

		inputtedText := r.FormValue("inputtedText")
		font := r.FormValue("font")
		templates.ExecuteTemplate(w, "index.html", askii(inputtedText, font))

	}

}
