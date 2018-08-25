package main

import (
	"fmt"
)

// declare 3 package level variables
var x int
var y string
var z bool

func main() {
	fmt.Println("Printing out the 'Zero Values', for our three package level variables:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
