package main

import "fmt"

func main() {
	compositeLiteral := [5]int{1, 2, 3, 4, 5}
	for i, v := range compositeLiteral {
		fmt.Printf("index = %v, value= %v\n", i, v)
	}

	fmt.Printf("%T\n", compositeLiteral)
}
