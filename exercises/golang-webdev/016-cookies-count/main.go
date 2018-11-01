package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./template/*.gohtml"))
}

func main() {
	// Routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Start Server
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	// increment cookie
	c := incrementPageView(req)
	// set cookie
	http.SetCookie(w, c)

	// render page
	tpl.ExecuteTemplate(w, "index.gohtml", c)
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	// increment cookie
	c := incrementPageView(req)
	// set cookie
	http.SetCookie(w, c)

	// render page
	tpl.ExecuteTemplate(w, "about.gohtml", c)
}

func incrementPageView(req *http.Request) *http.Cookie {
	// check if our cookie exists.
	c, err := req.Cookie("some-cookie")
	if err != nil {
		c = &http.Cookie{
			Name:  "some-cookie",
			Value: "1",
			Path:  "/",
		}
		return c
	}

	// convert to an integer
	i, err := strconv.Atoi(c.Value)
	if err != nil {
		return c
	}

	// increment value and reset it.
	i++
	c.Value = strconv.Itoa(i)

	return c
}
