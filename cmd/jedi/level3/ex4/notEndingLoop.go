package main

import "fmt"

func main() {
	age := 1980
	for {
		if age <= 2019 {
			fmt.Println(age)
			age++
		} else {
			break
		}
	}
}
