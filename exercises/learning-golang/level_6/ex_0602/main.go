/*
create a func with the identifier foo that
	takes in a variadic parameter of type int
	pass in a value of type []int into your func (unfurl the []int)
	returns the sum of all values of type int passed in
create a func with the identifier bar that
	takes in a parameter of type []int
	returns the sum of all values of type int passed in
*/
package main

import "fmt"

func main() {
	// tmp :=
	fmt.Println("calling foo(),", foo([]int{1, 2, 3, 4, 5}...))
	fmt.Println("calling bar(),", bar([]int{1, 20, 3, 4, 5}))
}

func foo(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func bar(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
