package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hi, I'm ", p.name)
}

type man struct {
	age int
}

func (m man) speak() {
	fmt.Println("I'm", m.age, "years old!")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	/*Receivers       Values
	-----------------------------------------------
		(t T)           T and *T
		(t *T)          *T
	*/

	p := person{name: "Bob"}
	saySomething(&p) // it must be pointer to p because speak method implementation has pointer receiver!

	m := man{age: 30}
	saySomething(m)
	//or
	saySomething(&m)
}
