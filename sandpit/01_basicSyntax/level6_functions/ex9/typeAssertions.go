package main

import "fmt"

//https://tour.golang.org/methods/15
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	switch i.(type) {
	case string:
		fmt.Printf("string type")
	}

	f = i.(float64) // panic
	fmt.Println(f)
}
