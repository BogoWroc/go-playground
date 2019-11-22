package main

import "fmt"

func main() {
	/*
		Deferred functions are executed before a function returns, even in the case of panic!
		recovery() function will stop panic and the program will continue running.
	*/
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	panic("I forgot my towel!!!")
}
