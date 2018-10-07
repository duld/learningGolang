/*
Using iota, create 4 constants for the last 4 years.
Print the constant values.
*/
package main

import "fmt"

const (
	yearBase = 2014 // iota still increments
	yearA    = yearBase + iota
	yearB    = yearBase + iota
	yearC    = yearBase + iota
	yearD    = yearBase + iota
)

func main() {
	fmt.Println(yearA)
	fmt.Println(yearB)
	fmt.Println(yearC)
	fmt.Println(yearD)
}
