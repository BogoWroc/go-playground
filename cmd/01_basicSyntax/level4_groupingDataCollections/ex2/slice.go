package main

import . "fmt"

func main() {
	compositeLiteral := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range compositeLiteral {
		Printf("index = %v, value= %v\n", i, v)
	}

	Printf("%T\n", compositeLiteral)
}
