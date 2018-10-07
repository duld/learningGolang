// Create and use an anonymous struct
package main

import "fmt"

func main() {
	myData := struct { // this describes the shape of our data structure
		uid      map[string]string
		dataList []string
		_rev     int
	}{ // now lets initialize it
		uid:      make(map[string]string),
		dataList: make([]string, 4),
		_rev:     23,
	}

	// add user ids and usernames to our data structure
	myData.uid["0123bvxAzk"] = "Jeromy Bobbishu"
	myData.uid["76abAzkdZf"] = "Cassey Snafflu"

	// create our seed values
	dataListSet := []string{"user ids", "user names", "location", "comments"}

	// populate myData.dataList with our seed values
	for i, v := range dataListSet {
		myData.dataList[i] = v
	}

	// print out the data
	fmt.Println(myData)

	// print out the data in a nicer format.
	fmt.Println("User IDs:")
	for k, v := range myData.uid {
		fmt.Printf("\t%v: %v\n", k, v)
	}
	fmt.Println("\nThe types of data we are interested in")
	for _, v := range myData.dataList {
		fmt.Printf("\t%v\n", v)
	}
	fmt.Println("\nCurrent rev:\n\t", myData._rev)
}
