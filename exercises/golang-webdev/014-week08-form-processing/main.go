/*
- Create a form on an HTML page that lets you submit values.
- Use FormValue to capture the values.
- Show the values on a page.
*/
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type Person struct {
	Name, Occupation, Age string
}

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	// Routes
	http.HandleFunc("/", indexHandler)

	// Start Server
	fmt.Println("Starting the server on port 8000")
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// GET
	if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	}

	// POST
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		p := Person{
			Name:       r.FormValue("name"),
			Occupation: r.FormValue("occupation"),
			Age:        r.FormValue("age"),
		}

		tpl.ExecuteTemplate(w, "index.gohtml", p)
	}
}
