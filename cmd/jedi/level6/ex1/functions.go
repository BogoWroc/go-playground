package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 10
}

func bar() (int, string) {
	return 5, "bla"
}
