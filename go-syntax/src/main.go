package main

import (
	"capstone"
)

func main() {
	capstone.PrintTable(capstone.CelsiusTableHeader, capstone.CelsiusTableRow)
	capstone.PrintTable(capstone.FahrenheitTableHeader, capstone.FahrenheitTableRow)
}
