/*
Fix the race condition you created in exercise #4 by using package atomic
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

// var mx sync.Mutex

func main() {
	r := rand.Intn(30)
	var total int64
	loopLimit := 200
	wg.Add(loopLimit)

	for i := 0; i < loopLimit; i++ {
		go func() {
			// mx.Lock()
			atomic.AddInt64(&total, 1)
			time.Sleep(time.Duration(r) * time.Millisecond)
			atomic.LoadInt64(&total)
			// mx.Unlock()
			wg.Done()
		}()
	}

	// wait for our go routines to finish
	wg.Wait()
	fmt.Println(total)
}
