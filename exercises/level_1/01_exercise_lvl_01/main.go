package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	// single line
	fmt.Println(x, y, z)

	// construct a line,
	fmt.Printf("%v ", x)
	fmt.Printf("%v ", y)
	fmt.Printf("%v\n", z)

	// give each one a line
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
