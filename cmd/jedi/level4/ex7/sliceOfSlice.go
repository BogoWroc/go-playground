package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	data := [][]string{
		xs1,
		xs2,
	}

	fmt.Println(data)

}
