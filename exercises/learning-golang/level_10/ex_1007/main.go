/*
write a program that
 - launches 10 goroutines
 - each goroutine adds 10 numbers to a channel
 - pull the numbers off the channel and print them
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	// send values - blocks due to wg
	go sendNNumsTimesTen(10, c)
	// read values
	readNums(c)
}

func sendNNumsTimesTen(n int, c chan int) {
	wg.Add(n)
	for i := 0; i < 10; i++ {
		go func() {
			sendNums(n, c)
		}()
	}
	wg.Wait()
	close(c)
}

func sendNums(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	wg.Done()
}

func readNums(c chan int) {
	for val := range c {
		fmt.Println(val)
	}
}
