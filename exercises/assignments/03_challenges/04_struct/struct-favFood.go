/*
Take the STRUCT “person” in the previous exercise and add a field called
“favFood” that stores a slice of string; for the variable “p1” use a composite
literal to add values to the field favFood; print out the values in favFood;
also print out the values for “favFood” by using a for range loop
*/
package main

import (
	"fmt"
)

// Person : stores first and last name of a person.
type Person struct {
	fName, lName string
	favFood      []string
}

func main() {
	// Using Struct person, create a value of type person and assign it to a
	// variable with the identifier p1
	p1 := Person{
		fName: "Bill",
		lName: "Krawschache",
		favFood: []string{
			"ribs",
			"turkey",
			"bacon",
		},
	}

	// Print out the values in favFood
	fmt.Println("Favorite foods slice:")
	fmt.Println(p1.favFood)

	// also print out the values for “favFood” by using a for range loop
	fmt.Println("\nFavorite foods:")
	for _, v := range p1.favFood {
		fmt.Println(v)
	}
}
