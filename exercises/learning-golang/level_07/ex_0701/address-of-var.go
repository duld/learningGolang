/*
Create a value and assign it to a variable
Print the address of that value
*/
package main

import (
	"fmt"
)

func main() {
	myVar := "Herro Dere!"
	fmt.Printf("I have created a variable 'myVar' with value: %v\n", myVar)
	fmt.Println("It lives at address: ", &myVar)
}
