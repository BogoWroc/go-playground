package main

import (
	"fmt"
	"os"
)

func main() {
	// go run programArguments.go one two

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
}
