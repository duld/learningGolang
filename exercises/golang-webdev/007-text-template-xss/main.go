// demonstrate text/template vs html/template for XSS
package main

import (
	"fmt"
	tplhtml "html/template"
	"log"
	"os"
	tpltxt "text/template"
)

var tplTxt *tpltxt.Template
var tplHTML *tplhtml.Template

type Post struct {
	Title string
	Body  string
}

// load the templates
func init() {
	// parse the templates directory for template files
	tplTxt = tpltxt.Must(tpltxt.ParseGlob("templates/*.gohtml"))
	tplHTML = tplhtml.Must(tplhtml.ParseGlob("templates/*.gohtml"))
}

func main() {
	var err error

	newPost := Post{
		Title: "First Post!",
		Body:  `<script>\nalert(\"Don't mind me, just encrypting your harddrive.....\")</script>`,
	}

	fmt.Println("------------")
	fmt.Println("--- TEXT ---")
	err = tplTxt.ExecuteTemplate(os.Stdout, "index.gohtml", newPost)
	if err != nil {
		log.Fatal("There was a problem reading the template from disk.", err)
	}

	fmt.Println()
	fmt.Println("------------")
	fmt.Println("--- HTML ---")
	err = tplHTML.ExecuteTemplate(os.Stdout, "index.gohtml", newPost)
	if err != nil {
		log.Fatal("There was a problem reading the template from disk.", err)
	}
}
