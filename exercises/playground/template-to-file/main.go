package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

// load in the templates
func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {

	// create two files
	// f, err := os.Create("authors.html")
	// if err != nil {
	// 	log.Fatal("There was a problem creating the authors file", err)
	// }
	// defer f.Close()

	// f2, err := os.Create("books.html")
	// if err != nil {
	// 	log.Fatal("There was a problem creating the books file", err)
	// }
	// defer f2.Close()

	// // write to the files
	// err = tpl.ExecuteTemplate(f, "authors.gohtml", nil)
	// if err != nil {
	// 	log.Fatal("Problem creating html file from template", err)
	// }

	// err = tpl.ExecuteTemplate(f2, "books.gohtml", nil)
	// if err != nil {
	// 	log.Fatal("Problem creating html file from template", err)
	// }
	xtemps := []string{"authors", "books"}
	for _, v := range xtemps {
		templateToFile(v)
	}
}

func templateToFile(templateName string) {
	f, err := os.Create(templateName + ".html")
	if err != nil {
		log.Fatal(fmt.Sprintf("There was a problem creating the %v file from template", templateName), err)
	}
	defer f.Close()

	err = tpl.ExecuteTemplate(f, templateName+".gohtml", nil)
	if err != nil {
		log.Fatal("Problem creating html file from template", err)
	}
}
