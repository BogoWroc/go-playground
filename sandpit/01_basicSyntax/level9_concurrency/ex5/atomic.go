package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func incCounter(nr int) {
	atomic.AddInt64(&counter, 1)
	fmt.Println("[Goroutine ", nr, "] counter = ", atomic.LoadInt64(&counter))
	wg.Done()
}

// to check if application has problem with race condition you should run application with flag -race
// go run -race atomic.go
func main() {
	fmt.Println("Start main")
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go incCounter(i)
	}

	fmt.Println("Wait for goroutines ...")
	wg.Wait()
	fmt.Println("Counter value is ", counter)
	fmt.Println("Done.")
}
