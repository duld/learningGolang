// parse n number of template files using the
// template.parseglob() method

package main

import (
	"html/template"
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
	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", books)
}
