package main

import "fmt"

var println = func(text string) {
	fmt.Println(text)
}

func main() {

	println("Some text to print")

	f := func() {
		println("Say hello ;)")
	}

	f()

	func(text string) { println(text) }("My name is Bob!")

}
