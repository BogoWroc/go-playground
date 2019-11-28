package main

import "fmt"

func main() {
	switchCase1()
	switchCase2()
	switchCase3()

}

func switchCase1() {
	name := "Bob"
	switch name {
	case "Bob", "Kev":
		fmt.Println("Male names")
	case "Alice", "Marta":
		fmt.Println("Female names")
	default:
		fmt.Println("Unknown sex")
	}
}

func switchCase2() {
	i := 5
	switch {
	case i > 4: // only first matched case will be executed!
		fmt.Println("i is greater than 4")
	case i > 3:
		fmt.Println("i is greater than 3")
	case i < 2:
		fmt.Println("i is lower than 2")
	}
}

func switchCase3() {
	i := 5
	switch {
	case i > 4:
		fmt.Println("i is greater than 4")
		fallthrough // if we want to process next cases we must use fallthrough, but this is not recommended
	case i > 3:
		fmt.Println("i is greater than 3")
	case i < 2:
		fmt.Println("i is lower than 2")
	}
}
