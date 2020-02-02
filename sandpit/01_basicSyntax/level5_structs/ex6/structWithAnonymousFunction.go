package main

import "fmt"

type Person struct {
	Name string
	RunE func(p *Person) string
}

func main() {
	p := Person{
		Name: "Bob",
		RunE: func(p *Person) string {
			return p.Name
		},
	}

	fmt.Println(p.RunE(&p))
}
