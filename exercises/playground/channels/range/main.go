package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// send
	go foo(c) // pass the channel to foo
	// channels block, so we have to set and read values from a channel on different
	// go processes.

	// receive
	for cv := range c {
		bar(cv)
	}

	fmt.Println("About to exit")
}

func foo(c chan int) {
	for i := 0; i < 70; i++ {
		c <- i
	}
	fmt.Println("The Channel has been popuated.")
	close(c)
}

func bar(c int) {
	fmt.Println(c)
}

// func foo(c chan int) {
// 	c <- 34
// 	close(c)
// }

// func bar(c chan int) {
// 	fmt.Println(<-c)
// }
