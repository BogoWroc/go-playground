package main

import "fmt"

func main() {
	func(value string) {
		fmt.Println(value)
	}("bla")
}
