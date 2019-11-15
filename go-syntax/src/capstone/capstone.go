package capstone

import "fmt"

type celsius float64
type fahrenheit float64

func (c celsius) toFahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func PrintTable(tableHeader func(), tableRow func(v celsius)) {
	fmt.Println("=====================")
	tableHeader()
	fmt.Println("=====================")
	for i := -40.0; i <= 100.00; i += 5 {
		c := celsius(i)
		tableRow(c)
	}
	fmt.Println("=====================")
}

func CelsiusTableRow(c celsius) {
	fmt.Printf("| %5.1f  |  %5.1f   |\n", c, c.toFahrenheit())
}

func CelsiusTableHeader() {
	fmt.Println("| C      |  F       |")
}

func FahrenheitTableRow(c celsius) {
	fmt.Printf("| %5.1f  |  %5.1f   |\n", c.toFahrenheit(), c)
}

func FahrenheitTableHeader() {
	fmt.Println("| F      |  C       |")
}
