package main

import (
	"fmt"
	"math/rand"
	"time"
)

var spaceLines = []string{"Space Adventures", "SpaceX", "Virgin Galactic"}
var tripTypes = []string{"One-way", "Round-trip"}

const minPrice = 36
const maxPrice = 50

func main() {

	printHeader()
	for i := 0; i < 10; i++ {
		spaceLine := randomGet(spaceLines)
		tripType := randomGet(tripTypes)
		currentDate := getRandomCurrentDate()
		departureInDays := calculateDepartureInDays(currentDate)
		speed := rand.Intn(15) + 16
		price := calculatePrice(speed, tripType)

		printTicket(spaceLine, tripType, departureInDays, price)
	}
}

func printHeader() {
	fmt.Println("Spaceline        Days Trip type  Price")
	fmt.Println("======================================")
}

func randomGet(data []string) string {
	length := len(data)
	return data[rand.Intn(length)]
}

func getRandomCurrentDate() time.Time {
	return Date(2019, rand.Intn(12)+1, rand.Intn(31)+1)
}

func calculateDepartureInDays(currentDate time.Time) int {

	departure := Date(2020, 10, 13)
	days := departure.Sub(currentDate).Hours() / 24
	return int(days)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func calculatePrice(speed int, tripType string) int {
	value := minPrice

	if speed > 16 && speed < 30 {
		value = rand.Intn(maxPrice-minPrice+1) + minPrice
	} else {
		value = maxPrice
	}

	if tripType == "Round-trip" {
		return value * 2
	} else {
		return value
	}

}

func printTicket(spaceLine string, tripType string, departureInDays int, price int) {
	fmt.Printf("%-16v %4v %-10v $%4v\n",
		spaceLine,
		departureInDays,
		tripType,
		price)
}
