/*
create a program that allows you to pass "James" into a template.
Print your template to std.Out
*/

package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

type Person struct {
	Name string
}

func init() {
	// load all the template files
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	myData := Person{"James"}
	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", myData)
}
