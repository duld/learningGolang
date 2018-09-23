/*
Fix the race condition you created in the previous exercise by using a mutex
- it makes sense to remove runtime.Gosched()
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mx sync.Mutex

func main() {
	r := rand.Intn(30)
	total := 0
	loopLimit := 200
	wg.Add(loopLimit)

	for i := 0; i < loopLimit; i++ {
		go func() {
			mx.Lock()
			loc := total
			time.Sleep(time.Duration(r) * time.Millisecond)
			loc++
			total = loc
			mx.Unlock()
			wg.Done()
		}()
	}

	// wait for our go routines to finish
	wg.Wait()
	fmt.Println(total)
}
