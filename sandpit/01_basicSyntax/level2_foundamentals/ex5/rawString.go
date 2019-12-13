package main

import "fmt"

func main() {
	s := "Some text"
	fmt.Println(s)

	rawString := `Some 
text
	to display`

	fmt.Println(rawString)
}
