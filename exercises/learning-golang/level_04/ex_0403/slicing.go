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
	nums := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// create a series of slices based on the values of the nums slice.
	numsA := nums[:5]
	numsB := nums[5:]
	numsC := nums[2:7]
	numsD := nums[1:6]

	// print out the new slices
	fmt.Println(numsA)
	fmt.Println(numsB)
	fmt.Println(numsC)
	fmt.Println(numsD)
}
