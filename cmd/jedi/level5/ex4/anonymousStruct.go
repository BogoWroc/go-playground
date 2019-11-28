package main

import "fmt"

func main() {
	p := struct {
		name string
		age  int
	}{
		name: "Bob",
		age:  55,
	}

	fmt.Println(p)
}
