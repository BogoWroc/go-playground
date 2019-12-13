package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

func incCounter(nr int) {
	mutex.Lock()
	counter = counter + 1
	fmt.Println("[Goroutine ", nr, "] counter = ", counter)
	mutex.Unlock()
	wg.Done()
}

// to check if application has problem with race condition you should run application with flag -race
// go run -race mutex.go
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
