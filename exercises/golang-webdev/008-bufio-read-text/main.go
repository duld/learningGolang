// use bufio.NewScanner to read some text
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var err error
	// open the text file
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("There was a problem opening the file", err)
	}

	line := 0 // keep track of our line number.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		fmt.Printf("line # %v: %v\n", line, scanner.Text())
		line++
	}
	if err = scanner.Err(); err != nil {
		log.Fatal("Problem reading the file", err)
	}
}

/*
bufio
- wraps io.Reader and io.Writer to create another object
*/
