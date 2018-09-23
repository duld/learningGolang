package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutimes:", runtime.NumGoroutine())

	r := rand.Intn(20)
	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			time.Sleep(time.Duration(r) * time.Millisecond)
			// runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutimes:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutimes:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
