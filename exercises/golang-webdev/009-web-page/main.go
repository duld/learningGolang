/*
Create a webpage that is served each day of the week. Include the following in each solution:

// Server //
- http.ListenAndServe
- http.HandleFunc
- func(http.ResponseWriter, *http.Request)

// Templates //
- var tpl *template.Template
- func init
- template.Must
- template.ParseGlob
- tpl.ExecuteTemplate
*/
package main

// func(w http.ResponseWriter, r *http.Request)

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	// parse our gohtml files
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {

	// Declare Routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	// start the server - listening for requests on port 8000
	fmt.Println("Starting server on port: 8000")
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		io.WriteString(w, "WHOOPS!")
	}
	// io.WriteString(w, "Hi")
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		io.WriteString(w, "WHOOPS!")
	}
	// io.WriteString(w, "About")
}
