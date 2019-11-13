package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggyBank := 0.0
	var money = []float64{0.05, 0.10, 0.25}

	for {
		piggyBank += money[rand.Intn(3)]
		if piggyBank >= 20.00 {
			break
		}
	}

	fmt.Printf("$%5.2f", piggyBank)
}
