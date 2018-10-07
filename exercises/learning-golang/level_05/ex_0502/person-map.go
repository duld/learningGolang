/*
Take the code from the previous exercise, then store the values of type person
in a map with the key of last name. Access each value in the map. Print out
the values, ranging over the slice.
*/
package main

import "fmt"

// Person : a representation of a person
type Person struct {
	first, last string
	favIceCream []string
}

func main() {
	jm := Person{
		first:       "Jeffrey",
		last:        "McCrea",
		favIceCream: []string{"Chocolate", "Cherry"},
	}

	kf := Person{
		first:       "Kesia",
		last:        "Flora",
		favIceCream: []string{"Vanilla", "Strawberry"},
	}

	// store the Person structs inside of a map
	ppl := map[string]Person{
		"mccrea": jm,
		"flora":  kf,
	}

	// Access each value in the map
	// -- Print out the values
	// -- range over the slice - favIceCream
	fmt.Println(ppl["mccrea"].first, ppl["mccrea"].last)
	for _, v := range ppl["mccrea"].favIceCream {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Println(ppl["flora"].first, ppl["flora"].last)
	for _, v := range ppl["flora"].favIceCream {
		fmt.Printf("\t%v\n", v)
	}
}
