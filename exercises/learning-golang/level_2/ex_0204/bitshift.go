/*
	assigns an int to a variable.
	prints that int in decimal, binary, and hex.
	shifts the bits of that it over 1 position to the left,
		and assigns that to a variable.
	prints that variable in decimal, binary and hex.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pre Bitshift")
	a := 45
	fmt.Printf("%d, %b, %x\n", a, a, a)

	fmt.Println("\nPost Bitshift")
	a = a << 1
	fmt.Printf("%d, %b, %x\n", a, a, a)
}
