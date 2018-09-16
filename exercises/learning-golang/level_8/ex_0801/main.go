/*
Starting with this code, marshal the []user to JSON. There is a little bit of
a curve ball in this hands-on exercise - remember to ask yourself what you
need to do to EXPORT a value from a package.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string `json:"first"`
	Age   int    `json:"age"`
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	usersJSON, err := json.Marshal(users)

	if err != nil {
		fmt.Println("There was a problem converting ")
	}
	fmt.Println(string(usersJSON))
	fmt.Printf("%T\n", usersJSON)
}
