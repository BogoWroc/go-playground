package main

import "fmt"

func calculate(data int) func() int {
	return func() int {
		return data + 2
	}
}

func main() {
	fmt.Println(calculate(5)())
}
