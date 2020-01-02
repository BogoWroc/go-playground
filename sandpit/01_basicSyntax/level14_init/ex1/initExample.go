package main

import "fmt"

/*
Unlike the main() function that can only be declared once, the init() function can be declared multiple times throughout a package. However, multiple init()s can make it difficult to know which one has priority over the others. In this section, we will show how to maintain control over multiple init() statements.
*/
func init() {
	fmt.Println("ex1 1")
}

func init() {
	fmt.Println("ex1 2")
}

func init() {
	fmt.Println("ex1 3")
}

func main() {
	fmt.Printf("main")
}
