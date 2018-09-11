/*
Use the “defer” keyword to show that a deferred func runs
after the func containing it exits.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start of main!")
	defer foo()
	fmt.Println("Foo was just called!")
	fmt.Println("This is the end of main...")
}

func foo() {
	fmt.Println("foo is being run!")
}
