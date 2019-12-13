package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doSomething() {
	fmt.Println("Do something ...")
	time.Sleep(time.Second * 2)
	wg.Done()
}

func main() {
	fmt.Println("Start main")
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	wg.Add(2)
	go doSomething()
	go doSomething()

	fmt.Println("Waiting for goroutines ...")
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	fmt.Println("Done.")

}
