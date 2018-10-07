// Create a program that shows the "if statement" in action
package main

import (
	"fmt"
)

const (
	limit = 300
)

func main() {
	n := 28

	if n < limit && n > 201 {
		fmt.Println("We can keep going!")
	} else if n < 30 {
		fmt.Println("What have you been doing??")
	} else {
		fmt.Println("Lets take a break....")
	}
}
