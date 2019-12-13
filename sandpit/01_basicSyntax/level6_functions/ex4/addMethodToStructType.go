package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("%v %v - %d old", p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Bob",
		last:  "Kent",
		age:   99,
	}

	p.speak()
}
