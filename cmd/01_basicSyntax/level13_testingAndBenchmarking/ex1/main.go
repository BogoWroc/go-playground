package main

import (
	"fmt"
	"github.com/bogowroc/go-playground/cmd/01_basicSyntax/level13_testingAndBenchmarking/ex1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.HumanYearsInDogYears(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.HumanYearsInDogYearsTwo(20))
}
