/*
- Create a struct of type person and give it a field named "first" of type
	string
- attach a method attached to type person called "speak"; the method should
	print something out
- create a value of type person and assign it to a variable named "p1"; call
	method speak
- create a slice of int; populate it with values; range of the slice of int
	printing out the index and value
- create a map of type [string]int; populate it with values; range over the
	map printing out the key and value
*/
package main

import "fmt"

// Person : represents a simple person data struct
type Person struct {
	first string
}

func (p Person) speak() {
	fmt.Printf("Hello, my name is %v, nice to meet you.\n", p.first)
}

func main() {
	fmt.Println("\nType Person:")
	p1 := Person{first: "Jefrey"}

	p1.speak()

	// create a slice of int; populate it with values, then range over it
	// printing out the index and values
	fmt.Println("\nSlice of int:")
	nums := []int{1, 22, 45, 63, 76, 83, 11, 3}

	for i, v := range nums {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// create a map of type [string]int; populate it with values; range over the
	// map printing out the key and value
	fmt.Println("\nMap:")
	myMap := map[string]int{
		"Kansas":     0,
		"Arkansas":   0,
		"Mexico":     1,
		"New-Mexico": 1,
		"Texas":      1,
		"Nevada":     1,
		"Arizona":    1,
		"Oregon":     1,
		"Montana":    1,
	}

	for k, v := range myMap {
		fmt.Println(k, v)
	}
}
