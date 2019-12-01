package main

import (
	"fmt"
	"time"
)

func main() {
	// create a channel to communicate app with gorutine (thread/process)
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c) // start thread/process
	}

	fmt.Println("### Main app is waiting for started gorutines ")
	timeout := time.After(5 * time.Second) // create channel on which event will be send when timeout occurs
	for i := 0; i < 5; i++ {
		select { // select waits until one case is ready and then runs it and its associated case statement.
		case gopherID := <-c: // wait for gorutine
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}

func sleepyGopher(id int, c chan int) {
	fmt.Println("... ", id, " go sleeping ...")
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id
}
