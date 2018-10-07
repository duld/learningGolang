/*
* Using a COMPOSITE LITERAL:
  - create a SLICE of TYPE int
  - assign 10 VALUES
* Range over the slice and print the values out.
* Using format printing
	- print out the TYPE of the slice
*/

package main

import "fmt"

func main() {
	// Create a slice using composite literal syntax.
	nums := []int{21, 33, 44, 55, 26, 87, 455, 1024, 1988, 2049}

	// Range over the slice and print the values out.
	for _, n := range nums {
		fmt.Println(n)
	}

	fmt.Printf("%T\n", nums)
}
