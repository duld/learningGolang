/*
Serve the files in the "starting-files" folder
To get your images to serve, use only this:

fs := http.FileServer(http.Dir("public"))

Hint: look to see what type FileServer
returns, then think it through.
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
	fs := http.StripPrefix("/", http.FileServer(http.Dir("./public")))
	http.Handle("/pics/", fs)

	http.HandleFunc("/", homeHandler)

	// start server
	fmt.Println("Server up and running on port 8000")
	http.ListenAndServe(":8000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
