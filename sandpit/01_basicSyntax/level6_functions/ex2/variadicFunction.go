package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	fmt.Println(foo(data...))
	fmt.Println(bar(foo, data))
}

func foo(e ...int) (sum int) {
	for _, v := range e {
		sum += v
	}

	return sum
}

func bar(f func(e ...int) int, data []int) int {
	return f(data...)
}
