/*
Starting with this code, sort the []user by

age
last
Also sort each []string “Sayings” for each user

print everything in a way that is pleasant
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u user) PrintInfo() {
	fmt.Printf("%v %v %v\n", u.First, u.Last, u.Age)
	for i, v := range u.Sayings {
		fmt.Printf("\t%v %v\n", i, v)
	}
}

func printUsers(users []user) {
	for _, usrv := range users {
		usrv.PrintInfo()
	}
}

func sortUserSayings(users []user) {
	for _, usrv := range users {
		usrv.sortSayings()
	}
}

func (u user) sortSayings() {
	sort.Strings(u.Sayings)
}

// ByAge an interface to help sort our users by their age
type ByAge []user

func (ba ByAge) Len() int           { return len(ba) }
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }

// ByLast an interface to help sort our users by their last name
type ByLast []user

func (bl ByLast) Len() int           { return len(bl) }
func (bl ByLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }
func (bl ByLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }

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

	fmt.Println("\nUsers pre-sorted:")
	fmt.Println(users)

	// your code goes here

	// AGE
	fmt.Println("\nSort the users by Age:")
	sort.Sort(ByAge(users))
	printUsers(users)

	// LAST
	fmt.Println("\nSort the users by Last Name:")
	sort.Sort(ByLast(users))
	printUsers(users)

	// SAYINGS
	fmt.Println("\nSort the user sayings:")
	sortUserSayings(users)
	printUsers(users)
}
