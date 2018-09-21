// parse n number of template files using the
// template.parseglob() method

package main

import (
	"html/template"
	"log"
	"os"
)

type Bookshelf struct {
	Books []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	books := Bookshelf{
		Books: []string{"The Great Gatsby",
			"War and Peace",
			"Great Expectations"},
	}

	// execute index.html
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", books)
	if err != nil {
		log.Fatal("Sorry, our server is not so great.")
	}
}
