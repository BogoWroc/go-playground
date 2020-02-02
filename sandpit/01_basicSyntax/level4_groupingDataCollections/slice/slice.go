package main

import "fmt"

func main() {
	slice := []string{"1", "2", "3", "4", "5"}

	fmt.Println(slice[:])
	fmt.Println(slice[1:])
	fmt.Println(slice[:3])

	// remove element from slice
	newSlice := append(slice[:2], slice[3:]...)
	fmt.Println(newSlice)
}
