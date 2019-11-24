package main

import "fmt"

func main() {
	var a int = 7
	var b *int = &a
	fmt.Println(a, b)
	a = 10
	fmt.Println(a, *b) // in this case * means dereference
	*b = 22
	fmt.Println(a, *b)
}
