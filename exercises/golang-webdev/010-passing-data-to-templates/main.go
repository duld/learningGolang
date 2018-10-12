/*
Create the following examples. Push them to github. Submit the url to the folder holding your solutions.

- pass a value of type int to a template and display the value in the template
- pass a value of type string to a template and display the value in the template
- pass a value of type bool to a template and display the value in the template
- pass a value of type []string to a template and display the values in the template
- pass a value of type map to a template and display the values in the template
- pass a value of type map to a template and display the keys and values in the template
- create a struct person. create some values of type person.
	add those values to a slice of person. pass the value of type []person to a template and
	display all of the values in the template
*/
package main

// func(w http.ResponseWriter, r *http.Request)

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type interestingData struct {
	NumBears       int
	ForestName     string
	PreparedToKill bool
	BearSpecies    []string
	BearAttacks    map[string]string
}

type person struct {
	Firstname, Lastname string
}

var tpl *template.Template

func init() {
	// parse our gohtml files
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

var bearData interestingData
var people []person

func main() {
	p1 := person{
		"Bill", "Brasky",
	}

	p2 := person{
		"Olive", "Oil",
	}

	p3 := person{
		"Susan", "McSusanFace",
	}

	people = []person{p1, p2, p3}

	// create our bear data
	bearData = interestingData{
		4,
		"Mauling National Park",
		true,
		[]string{"Grizzly", "Sloth Bear", "Brown Bear", "Kodiak"},
		map[string]string{
			"Bill": "BillBear",
		},
	}

	// Declare Routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/bears", bearHandler)

	// start the server - listening for requests on port 8000
	fmt.Println("Starting server on port: 8000")
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		io.WriteString(w, "WHOOPS!")
	}
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", people)
	if err != nil {
		io.WriteString(w, "WHOOPS!")
	}
}

func bearHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "bears.gohtml", bearData)
	if err != nil {
		io.WriteString(w, "WHOOPS!")
		go fmt.Println(err)
	}
}
