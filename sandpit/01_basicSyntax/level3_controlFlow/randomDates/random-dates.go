package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	for i := 0; i < 10; i++ {
		displayRandomDate()
	}
}

func displayRandomDate() {
	year := rand.Intn(2019) + 1
	month := rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		if year%2 == 0 {
			daysInMonth = 28
		} else {
			daysInMonth = 29
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}
