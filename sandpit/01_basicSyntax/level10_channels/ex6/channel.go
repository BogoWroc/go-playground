package main

import (
	"fmt"
)

func main() {

	ch := produceData()
	for v := range ch {
		fmt.Println(v)
	}

}

func produceData() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch

}
