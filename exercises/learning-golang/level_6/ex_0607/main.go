// Assign a func to a variable, then call that func

package main

import (
	"fmt"
)

func main() {
	// create an anonymous function that works as a counter.
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	// call counter 1
	counter()
	// call counter 2
	counter()
	fmt.Printf("%v\n", counter())
}
