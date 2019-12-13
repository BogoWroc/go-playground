package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var wg sync.WaitGroup

func incCounter(nr int) {
	counter = counter + 1
	runtime.Gosched()
	fmt.Println("[Goroutine ", nr, "] counter = ", counter)
	wg.Done()
}

// to check if application has problem with race condition you should run application with flag -race
// go run -race raceCondition.go
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
