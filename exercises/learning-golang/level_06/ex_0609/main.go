/*
A “callback” is when we pass a func into a func as an argument.
For this exercise, pass a func into a func as an argument
*/
package main

import (
	"fmt"
)

func main() {
	myCb := func(s string) {
		fmt.Println(s, "And thats all i have to say about that")
	}

	saySomething(myCb, "hey Jerk!")
}

// saySomething(cb func(string) string)

func saySomething(cb func(string), s string) {
	cb(s)
}
