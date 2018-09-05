/*
Initialize a SLICE of int using a composite literal;
print out the slice; range over the slice printing out
just the index; range over the slice printing out both
the index and the value
*/

package main

import (
	"fmt"
)

func main() {
	// initialize a slice of int using a composite literal
	nums := []int{93, 11, 82, 34, 17, 80, 500}

	// print out the slice
	fmt.Println("Slice:")
	fmt.Println(nums)

	// range over the slice printing just the index
	fmt.Println("Index values:")
	for i := range nums {
		fmt.Println(i)
	}

	// range over the slice printing out bot the index and the value
	fmt.Println("Index and Values:")
	for i, v := range nums {
		fmt.Println(i, v)
	}
}
