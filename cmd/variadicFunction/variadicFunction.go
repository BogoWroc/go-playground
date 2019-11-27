package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 9, 3, 4))
	fmt.Println(sum2(1, 9, 3, 4))
	fmt.Println(sum3(1, 9, 3, 4))
}

func sum(elements ...int) int {
	result := 0
	for _, e := range elements {
		result += e
	}

	return result
}

// or it could be written as below

func sum2(elements ...int) (result int) {
	for _, e := range elements {
		result += e
	}

	return
}

//or

func sum3(first int, elements ...int) int {
	result := first
	for _, e := range elements {
		result += e
	}

	return result
}
