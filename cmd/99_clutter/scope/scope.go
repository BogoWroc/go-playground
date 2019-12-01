package main

import "fmt"

var a = 0

func main() {
	var count = 0

	for count < 10 {
		fmt.Println(count)
		count++
	}

	for i := 0; i < 10; i++ {

	}
}

func test() string {
	return string(a)
}
