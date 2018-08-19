package main

import (
	"fmt"
)

type myUniqueType int

var x myUniqueType

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
