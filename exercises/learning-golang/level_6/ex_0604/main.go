/*
Create a user defined struct with
	the identifier “person”
	the fields:
		first
		last
		age
attach a method to type person with
	the identifier “speak”
	the method should have the person say their name and age
create a value of type person
call the method from the value of type person
*/

package main

import (
	"fmt"
)

// Person : represents a simple person object
type Person struct {
	first, last string
	age         int
}

func main() {
	p1 := Person{
		"Bill",
		"Karoughavahvah",
		34,
	}

	// make p1 speak!
	p1.speak()
}

func (p Person) speak() {
	fmt.Printf("Hello, nice to meet you. My name is %v %v.\n", p.first, p.last)
}
