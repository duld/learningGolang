/*
Demonstrate using 301, 303 and 307 http responses.
*/

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	// File-Server
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	http.Handle("/assets/", fs)

	// Routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/users", usersHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// start server
	fmt.Println("Server started on port :8000")
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "contact.gohtml", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// this will redirect to the contact route
	// tpl.ExecuteTemplate(w, "users.gohtml", nil)

	//
	if r.Method == "GET" {
		http.Redirect(w, r, "/contact", 307)

	} else if r.Method == "POST" { // Temporarily moved - API down.
		http.Redirect(w, r, "/about", 303)

	} else { // Permenantly Moved
		http.Redirect(w, r, "/", 301)
	}
}
