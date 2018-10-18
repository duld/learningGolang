/*
TODO:
fill templates
execute templates
serve files
*/

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// init

var tpl *template.Template

func init() {
	// load templates
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	// File-Server
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	http.Handle("/assets/", fs)

	// Routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	// start server
	fmt.Println("Server up and running on port:", 8000)
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}
