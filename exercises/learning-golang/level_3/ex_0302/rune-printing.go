/*
Print every rune code point of the uppercase alphabet three times

Your output should look like this:
1
	U+0041 'A'
	U+0041 'A'
	U+0041 'A'
2
	U+0042 'B'
	U+0042 'B'
	U+0042 'B'
*/
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 26; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i+64)
		}
	}
}
