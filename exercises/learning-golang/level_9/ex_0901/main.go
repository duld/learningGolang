/*
In addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Adding two go routines to our waitgroup")
	wg.Add(2)

	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	go handOffCustomer()
	go addSideItem()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	wg.Wait()
}

func handOffCustomer() {
	fmt.Println("Please pull forward to the 2nd window")
	wg.Done()
}

func addSideItem() {
	fmt.Println("Would you like fries with that?")
	wg.Done()
}
