// Create a program that shows the "if statement" in action
package main

import (
	"fmt"
)

const (
	limit = 300
)

func main() {
	n := 200

	if n < limit {
		fmt.Println("We can keep going!")
	} else {
		fmt.Println("Lets take a break....")
	}
}
