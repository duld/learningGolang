package main

import (
	"fmt"
)

func main() {
	j := foo("string value goes here. many thanks.")
	fmt.Println(j)
}

func foo(name string) int {
	i := 0
	fmt.Println(name)
	_, err := fmt.Scanf("%d", &i)

	if !(err == nil) {
		fmt.Println(err)
	}

	return i
}
