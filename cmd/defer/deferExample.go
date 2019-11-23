package main

import "fmt"

func main() {
	deferOrder()
	deferVariable()
}

func deferOrder() { //defer invocation executes in LIFO order Last In First Out! Additionally defer invocations
	// are executed at the end of method!
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func deferVariable() {
	a := "Bob"
	defer fmt.Println(a) // variable is set at the time of defer execution, not function execution!
	a = "Alice"
}
