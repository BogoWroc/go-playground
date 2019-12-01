package main

import "fmt"

const (
	a = iota
	b
	c
	d
)

const (
	e = iota + 2
	f
	g
	h
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println("###########")
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
