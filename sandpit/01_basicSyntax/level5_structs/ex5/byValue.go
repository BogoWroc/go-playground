package main

import "fmt"

type A struct {
	value int
}

type B struct {
	col []A
}

func NewB() B {
	return B{col: make([]A, 0)}
}

func (b *B) Add(a A) {
	a.value = 4
	b.col = append(b.col, a)
}
func add(m map[string]int) {
	m["test"] = 6
}

func add2(s []int) {
	x := append(s, 2)
	fmt.Println(x)
}
func main() {
	a := A{}

	a.value = 3

	b := NewB()
	b.Add(a)

	b.col = append(b.col, A{6})
	fmt.Println(b.col)
	fmt.Println(a)

	var x1 A
	fmt.Println(x1)

	var x2 [2]int
	fmt.Println(x2)

	var x3 []int
	add2(x3)
	fmt.Println(x3)

	var x4 map[string]int = make(map[string]int)
	add(x4)
	fmt.Println(x4)
}
