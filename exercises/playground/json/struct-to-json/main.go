package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	first, last string
	age         int
}

func main() {
	p1 := Person{"J", "M", 111}
	p2 := Person{"K", "F", 109}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("open the file....")
	// read json file
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("There was a problem opening the file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("Read the file data....")

	fi, _ := file.Stat()
	fileData := make([]byte, fi.Size())
	_, rerr := file.Read(fileData)
	if rerr != nil {
		fmt.Println("There was a problem reading the contents of the file")
		fmt.Println(rerr)
		return
	}

	fmt.Println(string(fileData))
	fmt.Println(len(fileData))
	fmt.Println("Convert slice of byte to interface....")
	// marshal data to an interface
	var mdata interface{}

	jerr := json.Unmarshal(fileData, &mdata)
	if jerr != nil {
		fmt.Println("There was a problem converting the slice byte to interface")
		fmt.Println(jerr)
	}

	// print out the results
	fmt.Println(mdata)
	fmt.Printf("%T\n", mdata)
}
