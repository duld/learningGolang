/*
Create a for loop using this syntax
	for condition {}
Have it print out th eyars you have been alive.
*/
package main

import "fmt"

func main() {
	birthYear := 1986
	curYear := 2018
	i := birthYear

	for i <= curYear {
		fmt.Println(i)
		i++
	}
	fmt.Println("Ain't dead yet!")
}
