// Create a program that uses a switch statement with no
// switch expression specified
package main

import (
	"fmt"
)

func main() {
	temp := 125
	goalTemp := 155

	switch {
	case temp > 100 && temp < 120:
		fmt.Println("Still need some time")
	case temp > 120 && temp < 140:
		fmt.Println("Getting closer!")
	case temp > 140 && temp < goalTemp:
		fmt.Println("Almost there!")
	case temp == goalTemp:
		fmt.Println("Take em out!")
	case temp > goalTemp:
		fmt.Println("NoooOOOOOooooOOoo!!!!")
	default:
		fmt.Println("Is the oven on?")
	}
}
