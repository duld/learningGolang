/*
Using the following operators, write expressions
and assign their values to variables:
	==
	<=
	>=
	!=
	<
	>
Now print each of the variables.
*/
package main

import (
	"fmt"
)

var (
	a bool
	b bool
	c bool
	d bool
	e bool
	f bool
)

func main() {
	a = 11 == 13 // false
	b = 13 <= 14 // true
	c = 55 >= 78 // false
	d = 77 != 77 // false
	e = 44 < 55  // true
	f = 76 > 12  // true

	fmt.Println(a, b, c, d, e, f)
}
