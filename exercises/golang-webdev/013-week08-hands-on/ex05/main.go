/*
Serve the files in the "starting-files" folder
To get your images to serve, use:

	func StripPrefix(prefix string, h Handler) Handler
	func FileServer(root FileSystem) Handler

Constraint: you are not allowed to change the route being used for images in
the template file
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
	fs := http.StripPrefix("/public", http.FileServer(http.Dir("./public")))
	http.Handle("/public/", fs)

	// Routes
	http.HandleFunc("/", indexHandler)

	// start server
	fmt.Println("starting server:")
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
