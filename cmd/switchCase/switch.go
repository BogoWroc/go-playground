package main

import "fmt"

func main() {
	simpleSwitch()
	switchWithInitializationSection()
	switchWithConditions()
	switchWithFallthrough()
	typeSwitch()
}

func typeSwitch() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	default:
		fmt.Println("unsupported type")

	}
}

func switchWithFallthrough() {
	i := 10
	switch {
	case i == 10:
		fmt.Println("Below 10")
		fallthrough
	case i > 10:
		fmt.Println("This line also be executed!")
	}
}

func switchWithConditions() {
	i := 10
	switch {
	case i < 10:
		fmt.Println("Below 10")
	case i >= 10:
		fmt.Println("Above or equals 10")
	}
}

func switchWithInitializationSection() {
	switch i := 2 + 3; i {
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("default")

	}
}

func simpleSwitch() {
	switch 1 {
	case 1, 2, 3:
		fmt.Println("1 or 2 or 3")
		fmt.Println("This line will also be executed")
	default:
		fmt.Println("default")
	}
}
