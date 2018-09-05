/*
Create a new type called person which is STRUCT that stores fName and lName
Using the STRUCT “person”, using a composite literal create a value of type
person and assign it to a variable with the identifier “p1”; print out “p1”;
print out just the field fName for “p1”
*/
package main

import (
	"fmt"
)

// Person : stores first and last name of a person.
type Person struct {
	fName, lName string
}

func main() {
	// Using Struct person, create a value of type person and assign it to a
	// variable with the identifier p1
	p1 := Person{
		fName: "Bill",
		lName: "Krawschache",
	}

	// Print out just the field fName for p1
	fmt.Println(p1.fName)

}
