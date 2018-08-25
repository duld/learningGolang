// Write a program that prints a number in decimal binary and hexidecimal
package main

import (
	"fmt"
)

func main() {
	num := 78
	fmt.Printf("Decimal: %d\n", num)
	fmt.Printf("Decimal: %b\n", num)
	fmt.Printf("Decimal: %x\n", num)
}
