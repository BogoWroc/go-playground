package main

import "fmt"

func main() {
	data := "some text"
	for i, c := range data {
		fmt.Printf("%d -> %c\n", i, c)
	}

	rawData := `some
				text
				in few lines`

	for _, c := range rawData {
		fmt.Printf("%c\n", c)

	}
}
