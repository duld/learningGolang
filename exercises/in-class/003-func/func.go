package main

import (
	"fmt"
)

func main() {
	fmt.Println("inside main")
	myFunc()
}

// func receiver identifer(parameters) returns {code}
func myFunc() {
	fmt.Println("this is coming from the Function call")
}
