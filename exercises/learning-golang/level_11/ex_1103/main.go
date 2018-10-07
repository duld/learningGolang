/*
code: https://play.golang.org/p/L5QWV8-p11
Create a struct “customErr” which implements the builtin.error interface.
Create a func “foo” that has a value of type error as a parameter. Create a
value of type “customErr” and pass it into “foo”. If you need a hint, here is one.
*/
package main

import (
	"fmt"
)

type customErr struct {
	msg string
}

func (ce customErr) Error() string {
	return ce.msg
}

func main() {
	c1 := customErr{"There was a problem when attempting to tie your shoes. Ask your father for help."}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
