package stringy

import "fmt"

// ToString - converts something to a string
func ToString(n interface{}) string {
	return fmt.Sprintf("%v", n)
}
