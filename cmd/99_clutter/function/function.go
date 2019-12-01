package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func writeToMethodParameter(s *string) {
	*s = "New value"
}

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%s %d", p.name, p.age)
}

func (p *person) setAge(age int) {
	p.age = age
}

func main() {
	fmt.Println(add(2, 3))

	p := person{
		name: "Bob",
		age:  60,
	}

	fmt.Println(p)
	p.setAge(24)
	fmt.Println(p)

	s := "text"
	writeToMethodParameter(&s)
	fmt.Println(s)
}
