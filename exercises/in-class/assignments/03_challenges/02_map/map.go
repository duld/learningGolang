/*
Initialize a MAP using a composite literal where the key is a string and the
value is an int; print out the map; range over the map printing out just the
key; range over the map printing out both the key and the value
*/

package main

import (
	"fmt"
)

func main() {
	// Initialize a map using a composite literal where the key is a string
	// and the value is an int
	myData := map[string]int{
		"la-to-austin":       1378,
		"la-to-vancouver":    1279,
		"la-to-grand-canyon": 486,
		"la-to-dennys":       0,
	}

	// Print out the map
	fmt.Println("Map:")
	fmt.Println(myData)

	// range over the map printing out just the key
	fmt.Println("\nKeys:")
	for k := range myData {
		fmt.Println(k)
	}

	// range over the map printing key/value pairs
	fmt.Println("\nKey/Values:")
	for k, v := range myData {
		fmt.Printf("%v: %v\n", k, v)
	}
}
