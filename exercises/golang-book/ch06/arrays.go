package main

import (
	"fmt"
)

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64
	/*
		for i := 0; i < len(x); i++ {
			total += x[i]
		}
	*/
	// loop over x
	for i, value := range x {
		total += value
		fmt.Println(i, value)
	}
	// print out the average of the values in the array.
	fmt.Println(total / float64(len(x)))

	// create a map[string]string
	var y = make(map[string]string)
	y["Name"] = "Jeff"
	y["Occupation"] = "Grunt"

	// loop over our map
	for key, value := range y {
		fmt.Println(key, value)
	}

}
