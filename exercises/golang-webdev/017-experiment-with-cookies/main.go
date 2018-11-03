/*
Experiment with the "path" attribute on cookies.

Trying setting a cookie when at this path "/dog/bowzer"
and see if you can access it at "/cat" and/or at "/dog/bowzer"

Trying setting a cookie when at this path "/" and see if you
can access it at "/cat" and/or "/dog/bowzer" and/or "/"
*/
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./template/*.gohtml"))
}

func main() {

	// Routes
	http.HandleFunc("/", indexHandler)               // should set a root level cookie
	http.HandleFunc("/dog/bowzer", dogbowzerHandler) // should set a cookie only readible at this path
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/cat", catHandler) // should set a cookie that is only readible at this path

	// Start Server
	fmt.Println("Server is starting on port :8000")
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	// set root level cookie
	xc := req.Cookies()
	if len(xc) == 0 {
		c := &http.Cookie{
			Name:  "root-level",
			Value: "howdy",
			Path:  "/",
		}
		http.SetCookie(w, c)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", xc)
}

func dogbowzerHandler(w http.ResponseWriter, req *http.Request) {
	// set dog/bowzer level cookie
	xc := req.Cookies()
	if len(xc) == 0 {
		c := &http.Cookie{
			Name:  "bowzer-level",
			Value: "A-Ruff...",
			Path:  "/dog/bowzer",
		}
		http.SetCookie(w, c)
	}

	tpl.ExecuteTemplate(w, "bowzer.gohtml", xc)
}

func dogHandler(w http.ResponseWriter, req *http.Request) {
	// dogs do nothing
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func catHandler(w http.ResponseWriter, req *http.Request) {
	// set cat level cookie
	xc := req.Cookies()
	if len(xc) == 0 {
		c := &http.Cookie{
			Name:  "cat-level",
			Value: "meow",
			Path:  "/cat",
		}
		http.SetCookie(w, c)
	}

	tpl.ExecuteTemplate(w, "cat.gohtml", xc)
}
