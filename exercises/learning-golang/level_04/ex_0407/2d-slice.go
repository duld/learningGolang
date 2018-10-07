/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
	- "James", "Bond", "Shaken, not stirred"
	- "Miss", "Moneypenny", "Helloooooo, James."
Range over the records, then range over the data in each record.
*/
package main

import "fmt"

func main() {
	bond := []string{"James", "Bond", "Shaken, not stirred"}
	penny := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	people := [][]string{bond, penny}

	for _, p := range people {
		fmt.Println("----------")
		for _, val := range p {
			fmt.Println(val)
		}
	}
}
