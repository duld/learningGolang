/*
Using a composite literal
	create an array which holds 5 values of type int
	assign values to each index position
Range over the array and print the values out
Using format printing
  print the type of the array
*/
package main

import "fmt"

func main() {
	// create an array which holds 5 values of type int
	nums := [5]int{23, 45, 87, 99, 100}

	// rang over the array and print the values out
	for _, v := range nums {
		fmt.Println(v)
	}

	// print the type of the array.
	fmt.Printf("%T\n", nums)
}
