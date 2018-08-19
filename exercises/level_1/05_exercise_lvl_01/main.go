package main

import (
	"fmt"
)

type myUniqueType int

var x myUniqueType
var y int

func main() {
	// output x
	fmt.Println("Printing custom type:")
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	// convert x to underlying type int
	y = int(x)
	fmt.Println("\nConverting custom type x->y:")
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
