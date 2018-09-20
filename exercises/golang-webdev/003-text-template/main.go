package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("one.txt", "two.txt")
	if err != nil {
		log.Fatal("Reading the files, didn't go over so well!", err)
	}

	// write to stdout.
	tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)
	tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)
}
