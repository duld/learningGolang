/*
Create a program that uses a switch statement with the switch
expression specified as a variable of TYPE string with the
IDENTIFIER “favSport”.
*/
package main

import (
	"fmt"
)

func main() {
	favSport := "running"

	switch favSport {
	case "swimming":
		fmt.Println("Lets go for a swim!")
	case "fencing":
		fmt.Println("Not in the face ok?")
	case "dancing":
		fmt.Println("How is this scored?")
	case "running":
		fmt.Println("Don't celebrate till you pass the finish line")
	default:
		fmt.Println("I'm not well versed in that sport.")
	}
}
