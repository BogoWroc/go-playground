package main

import "fmt"

func main() {
	ifWithInitializationSection()
}

func ifWithInitializationSection() {
	m := map[string]int{"Bob": 23, "Alice": 19}
	if age, ok := m["Bob"]; ok {
		fmt.Println("Bob has", age, "years old.")
	}
}
