/*
ListenAndServe on port 8080 of localhost
For the default route "/" Have a func called "foo" which writes
to the response "foo ran"

For the route "/dog/" Have a func called "dog" which parses a
template called "dog.gohtml" and writes to the response "

This is from dog
" and also shows a picture of a dog when the template is executed.
Use "http.ServeFile" to serve the file "dog.jpeg"
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
	// fileserver
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	http.Handle("/assets/", fs)

	// Routes
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dogHandler)

	// start server
	fmt.Println("Server up and running on port 8000")
	http.ListenAndServe(":8000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo ran")
}

func dogHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}
