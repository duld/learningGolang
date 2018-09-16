/*
Starting with this code, encode to JSON the []user sending the results to Stdout.
Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	// encode the users as JSON
	usersJSON, err := json.Marshal(users) // not needed after all....

	if err != nil {
		fmt.Println("There was a problem marshalling the data to json.")
	}

	// fmt.Println(string(usersJSON))
	enc := json.NewEncoder(os.Stdout)
	// write out as a string
	fmt.Println("\nOutput as a string:")
	err = enc.Encode(string(usersJSON))

	if err != nil {
		fmt.Println("Problem encoding json to writer")
	}

	fmt.Println("\nEncode without marshaling first:")
	err = enc.Encode(users)

	if err != nil {
		fmt.Println("Problem encoding json to writer")
	}
}
