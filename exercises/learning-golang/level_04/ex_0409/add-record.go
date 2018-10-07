/*
Using the code from the previous example, add a record to your map.
Now print the map out using the “range” loop
*/
package main

import "fmt"

func main() {
	peopleMap := make(map[string][]string)
	// seed the map
	peopleMap["bond_james"] = []string{
		`Shaken, not stirred`,
		`Martinis`,
		`Women`,
	}
	peopleMap["moneypenny_miss"] = []string{
		`James Bond`,
		`Literature`,
		`Computer Science`,
	}
	peopleMap["no_dr"] = []string{
		`Being evil`,
		`Ice cream`,
		`Sunsets`,
	}

	// adding a new record
	peopleMap["goldmember"] = []string{
		"Unfortunate shmelting accident",
		"Gold",
		"roller boogie",
	}

	for k, v := range peopleMap {
		fmt.Println(k)
		for i := 0; i < len(v); i++ {
			fmt.Printf("\t%v: %v\n", i, v[i])
		}
	}
}
