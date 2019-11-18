package main

import (
	"capstone"
	"chess"
	"fmt"
)

func main() {
	terraformExercise()
}

func terraformExercise() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	capstone.Planets(planets[3:4]).Terraform()
	capstone.Planets(planets[6:]).Terraform()
	fmt.Println(planets)

	planets = []string{"Venus", "Mars", "Jupiter"}
	newPlanets := capstone.Terraform("New", planets...)
	fmt.Println(newPlanets)

}

func chessExercise() {
	chess.DisplayBoard()
}

func capstoneExercise() {
	capstone.PrintTable(capstone.CelsiusTableHeader, capstone.CelsiusTableRow)
	capstone.PrintTable(capstone.FahrenheitTableHeader, capstone.FahrenheitTableRow)
}
