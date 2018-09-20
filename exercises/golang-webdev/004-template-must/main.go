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
	fmt.Println("fileA")
	err := tpl.ExecuteTemplate(os.Stdout, "fileA.gohtml", nil)
	if err != nil {
		panic(err)
	}

	// fileB
	fmt.Println("\n\nfileB")
	err = tpl.ExecuteTemplate(os.Stdout, "fileB.gohtml", nil)
	if err != nil {
		panic(err)
	}

	// fileC
	fmt.Println("\n\nfileC")
	err = tpl.ExecuteTemplate(os.Stdout, "fileC.gohtml", nil)
	if err != nil {
		panic(err)
	}
}
