package main

import "fmt"

type animal interface {
	Name() string
}

func New(n string) animal {
	return &dog{n}
}

type dog struct {
	name string
}

func (d *dog) Name() string {
	return d.name
}

func main() {
	a := New("Volta")
	fmt.Println(a.Name())
}
