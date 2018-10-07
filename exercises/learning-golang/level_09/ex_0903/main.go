/*
Using goroutines, create an incrementer program
  have a variable to hold the incrementer value
  launch a bunch of goroutines
  each goroutine should
    read the incrementer value
    store it in a new variable
		yield the processor with runtime.Gosched()
		increment the new variable
		write the value in the new variable back to the incrementer variable
use waitgroups to wait for all of your goroutines to finish
the above will create a race condition.
Prove that it is a race condition by using the -race flag
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	r := rand.Intn(30)
	total := 0
	loopLimit := 200
	wg.Add(loopLimit)

	for i := 0; i < loopLimit; i++ {
		go func() {
			loc := total
			time.Sleep(time.Duration(r) * time.Millisecond)
			loc++
			total = loc
			wg.Done()
		}()
	}

	// wait for our go routines to finish
	wg.Wait()
	fmt.Println(total)
}
