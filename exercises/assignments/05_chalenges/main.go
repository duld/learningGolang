package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	c := make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

// read the contents of a directory

// print the contents of a directory

// write a file to the directory

// read the file from the directory

// write over the file in the directory

// append data to the end of the file

// Write a json file

// read the json file back

// handle common errors

// write a simple sign up sheet program.
