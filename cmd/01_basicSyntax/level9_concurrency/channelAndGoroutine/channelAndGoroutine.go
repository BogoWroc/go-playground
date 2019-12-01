package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Some initialisation before main execution ... sweet :)")
}

func doSomething(value int) int {
	fmt.Println("[Goroutine] Do calculation ....")
	time.Sleep(3 * time.Second)
	return value + 1
}
func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println("Start waiting on response ...")
	fmt.Println("Done ... I received: ", <-ch)
	fmt.Println("End.")
}
