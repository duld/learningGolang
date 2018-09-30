package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)
	fmt.Println("exiting main...")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i&1 == 1 {
			o <- i
		} else {
			e <- i
		}
	}
	q <- 0 // exit
	defer close(e)
	defer close(o)
	defer close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Even:", v)
		case v := <-o:
			fmt.Println("Odd: ", v)
		case v := <-q:
			fmt.Println("Time to quit:", v)
			return // exit for loop
		}

	}
}
