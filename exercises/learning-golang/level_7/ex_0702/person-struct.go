/*
- create a person struct
- create a func called “changeMe” with a *person as a parameter
		- change the value stored at the *person address
- create a value of type person
- print out the value
- in func main
  - call “changeMe”
- in func main
  - print out the value
*/
package main

import (
	"fmt"
)

// Person : blah blah blah
type Person struct {
	First, Last string
}

func changeMe(p *Person, first string) {
	(*p).First = first // dereference the pointer, THEN access the field 'First'
}

func main() {
	jm := Person{"J", "M"}
	fmt.Println("This is the current value of my person: ", jm)
	fmt.Println("Lets update the name to something better!")

	changeMe(&jm, "Billrwacka'tha")
	fmt.Println("Here is the updated value of person: ", jm)
}
