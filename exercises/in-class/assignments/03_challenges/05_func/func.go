/*
Continuing with the code from here, add a method to type “person” with the
identifier “walk”. Have the func print this string: <person’s first name> is
walking. Call the method.
*/

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"Miss", "Moneypenny", 24}
	fmt.Println(p1)
	p1.foo()

	p1.walk()
}

func (p person) foo() {
	fmt.Println("Hello from foo")
}

func (p person) walk() {
	fmt.Println(p.first, "is walking")
}

// func (receiver) identifier(parameters) (returns) {}
