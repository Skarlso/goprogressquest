package main

import (
	"html/template"
	"log"
	"net/http"
)

//Page represents a page object
type Page struct {
	Title string
	Body  template.HTML
}

var templates = template.Must(template.ParseGlob("view/*.html"))

// The main function which starts the rpg.
func main() {
	http.HandleFunc("/", frontPageHandler)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./view/css"))))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./view/img"))))
	log.Printf("Starting server to listen on port: 8989...")
	http.ListenAndServe(":8989", nil)
}

func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Redirecting to FrontPage:")
	p := &Page{Title: "Welcome Page", Body: ""}
	renderTemplate(w, "index", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	log.Printf("Rendering template: %s", tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
