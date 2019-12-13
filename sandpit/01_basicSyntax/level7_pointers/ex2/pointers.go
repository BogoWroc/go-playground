package main

import "fmt"

type person struct {
	name string
}

func (p *person) changeMe(name string) {
	p.name = name
}

func changeMe2(p *person) { // if we use pointer type we can only pass a value of pointer type
	p.name = "Miss"
	//or
	(*p).name = "Kris"
}

func display(p person) { // it accepts only values
	fmt.Println(p.name)
}
func main() {
	p := person{name: "Bob"}
	fmt.Println(p)
	p.changeMe("John")
	fmt.Println(p)

	changeMe2(&p) // pointer
	fmt.Println(p)

	display(p)
}
