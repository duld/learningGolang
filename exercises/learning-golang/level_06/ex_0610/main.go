/*
Closure is when we have “enclosed” the scope of a variable in some code block.
For this hands-on exercise, create a func which “encloses” the scope of a
variable:
*/
package main

import (
	"fmt"
)

func main() {

	b := addItUp()
	fmt.Println(b([]int{1, 2, 3, 4, 5}...))
}

func addItUp() func(...int) int {
	total := 0

	return func(nums ...int) int {
		for _, v := range nums {
			total += v
		}
		return total
	}
}
