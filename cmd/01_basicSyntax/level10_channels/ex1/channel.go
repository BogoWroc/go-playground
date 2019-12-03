package main

import (
	"fmt"
)

func main() {
	channel()
	bufferedChannel()
}

func channel() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func bufferedChannel() {
	c := make(chan int, 1)

	c <- 10

	fmt.Println(<-c)
}
