// using the code linked as a base, pull off the values using a select
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case cVal, ok := <-c:
			if ok {
				fmt.Println(cVal)
			} else {
				fmt.Println("c is closed")
				return
			}
		case qVal, ok := <-q:
			if ok {
				fmt.Println(qVal)
			} else {
				fmt.Println("q is closed")
			}
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 23
		close(c)
	}()

	return c
}
