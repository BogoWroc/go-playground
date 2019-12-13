package main

import "fmt"

//in Go ONLY slice and map work at internal pointers!

func main() {
	// all assignment operations in Go are copy operations!

	sliceExample()
	mapExample()
}

func sliceExample() {
	a := []int{1, 2, 3}
	b := a
	b[1] = 10
	fmt.Println(a, b) // all slices will have the same value 1,10,3
	// slice is a projection to array. It has three fields: len, cap and pointer to array
}

func mapExample() {
	m := map[string]int{"bob": 23, "alice": 20}
	k := m
	k["bob"] = 33
	fmt.Println(m, k) // both map will have the same value
}
