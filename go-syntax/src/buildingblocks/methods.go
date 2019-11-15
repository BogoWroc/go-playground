package main

import "fmt"

type kelvin float64

// kelvinToCelsius converts °K to °C
func (k kelvin) celcius() celcius {
	return celcius(k - 273.15)
}

func (k kelvin) test(i int) {
	fmt.Println(i)
}

type celcius float64

func (c celcius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

type fahrenheit float64

func (k kelvin) fahrenheit() fahrenheit {
	return k.celcius().fahrenheit()
}

func main() {
	k := kelvin(294.0)
	k.test(100)
	celsius := k.celcius()
	fahrenheit := celsius.fahrenheit()
	fmt.Print(k, "° K is ", celsius, "° C is ", fahrenheit, "° F\n")
	fmt.Print(0, "° K is ", kelvin(0).fahrenheit(), "° F")
}
