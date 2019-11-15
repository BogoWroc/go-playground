package main

import (
	"fmt"
	"math/rand"
)

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

type rowSupplier func(row int) (string, string)

func drawTable(rows int, getRow rowSupplier) {

}

func main() {
	var sensor func() kelvin = fakeSensor

	fmt.Println(sensor())
	sensor = realSensor
	fmt.Println(sensor())
}
