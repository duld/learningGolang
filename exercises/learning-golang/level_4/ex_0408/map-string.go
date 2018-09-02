/*
Create a map with a key of TYPE string which is a person’s “last_first” name,
and a value of TYPE []string which stores their favorite things. Store three
records in your map. Print out all of the values, along with their index
position in the slice.
		`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
		`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
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

	for k, v := range peopleMap {
		fmt.Println(k)
		for i := 0; i < len(v); i++ {
			fmt.Printf("\t%v: %v\n", i, v[i])
		}
	}
}
