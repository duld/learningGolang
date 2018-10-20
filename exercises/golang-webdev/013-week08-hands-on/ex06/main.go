/*
Starting with the code in the "starting-files" folder:
wire this program up so that it works

ParseGlob in an init function

Use HandleFunc for each of the routes

Combine apply & applyProcess into one func called "apply"

Inside the func "apply", use this code to create the logic to respond
differently to a POST method and a GET method

if req.Method == http.MethodPost {
    		// code here
    		return
    	}
*/
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	// Routes
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	// http.HandleFunc("/applyProcess", applyProcess)

	// Start Server
	fmt.Println("Server up and running on port :8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {
	var err error
	// GET
	if req.Method == http.MethodGet {
		err = tpl.ExecuteTemplate(w, "apply.gohtml", nil)
		HandleError(w, err)
	}
	// POST
	if req.Method == http.MethodPost {
		err = tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		HandleError(w, err)
	}
}

func applyProcess(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	}

	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
