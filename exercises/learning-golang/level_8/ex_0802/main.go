/*
Starting with this code, unmarshal the JSON into a
Go data structure. Hint: https://mholt.github.io/json-to-go/
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First, Last string
	Age         int
	Sayings     []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var people []Person
	err := json.Unmarshal([]byte(s), &people)

	if err != nil {
		fmt.Println("There was a problem converting the json to a data structure")
	}

	fmt.Println("\nPost unmarshal")
	for _, v := range people {
		fmt.Println()
		fmt.Println(v.First, v.Last)
		fmt.Println(v.Age)
		for i, cph := range v.Sayings {
			fmt.Printf("\t%v. %v\n", i, cph)
		}
	}

}
