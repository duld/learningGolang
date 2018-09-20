/*
Create a program to parse three files into a VALUE of TYPE *template.Template

Use these funcs:
 - func init
 - template.Must
 - template.ParseGlob
*/
package main

import (
	"fmt"
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	// fileA
	err := tpl.ExecuteTemplate(os.Stdout, "fileA.gohtml", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n")

	// fileB
	err = tpl.ExecuteTemplate(os.Stdout, "fileB.gohtml", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n")

	// fileC
	err = tpl.ExecuteTemplate(os.Stdout, "fileC.gohtml", nil)
	if err != nil {
		panic(err)
	}
}
