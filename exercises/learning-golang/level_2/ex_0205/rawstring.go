// Create a variable of type string using a raw string literal.
// Then print the string.
package main

import "fmt"

func main() {
	myString := `This is a raw string ( ) where %d the rules ?!\n dont apply!
	 \t \t \r \n`
	fmt.Println(myString)
}
