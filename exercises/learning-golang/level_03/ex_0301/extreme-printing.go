// Print every number from 1 to 10,000
package main

import (
	"fmt"
)

func main() {
	const end = 10000
	for i := 1; i <= end; i++ {
		fmt.Println(i)
	}
}
