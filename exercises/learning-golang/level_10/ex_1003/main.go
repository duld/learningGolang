/*
Starting with this code, pull the values off the channel
using a for range loop
*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go gen(c)

	receive(c)

	fmt.Println("about to exit")
}

func gen(c chan int) {
	// c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
