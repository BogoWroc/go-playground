package main

import "fmt"

func main() {
	fmt.Println("start")
	startPanic()
	fmt.Println("end")
}

func startPanic() {
	fmt.Println("startPanic")
	defer func() { // defer always requires function invocation, not reference to function!
		if err := recover(); err != nil { // recovery handle panic event
			fmt.Println(err)
			//panic(err) // if we want to rethrow panic we need to uncomment this line
		}
	}()

	panic("something bad happened")
}
