package main

import (
	"capstone"
	"chess"
	"cmdline"
	"fmt"
	"time"
)

func main() {
	cmdlineArgsExercise()
}

func cmdlineArgsExercise() {
	cmdline.PrintArgs()
}

func embeddingExercise() {
	c := capstone.Car{
		Name:   "BMW",
		Engine: capstone.Engine{Capacity: 1.6},
	}
	c.Start()
}

func classesExercise() {
	spirit := capstone.NewLocation(capstone.Coordinate{14, 34, 6.2, 'S'}, capstone.Coordinate{175, 28, 21.5, 'E'})
	opportunity := capstone.NewLocation(capstone.Coordinate{1, 56, 46.3, 'S'}, capstone.Coordinate{354, 28, 24.2, 'E'})
	curiosity := capstone.NewLocation(capstone.Coordinate{4, 35, 22.2, 'S'}, capstone.Coordinate{137, 26, 30.12, 'E'})
	insight := capstone.NewLocation(capstone.Coordinate{4, 30, 0.0, 'N'}, capstone.Coordinate{135, 54, 0, 'E'})
	fmt.Println("Spirit", spirit)
	fmt.Println("Opportunity", opportunity)
	fmt.Println("Curiosity", curiosity)
	fmt.Println("InSight", insight)
}

func landingExercise() {
	capstone.PrintLandingSites()
}

func lifeExercise() {
	a, b := capstone.NewUniverse(), capstone.NewUniverse()
	a.Seed()
	for i := 0; i < 300; i++ {
		capstone.Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a
	}
}

func wordsExercise() {
	text := `As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far
			overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked.            
			Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was
			the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise
			there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a
			forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man.`
	frequency := capstone.Sentence(text).CountWords()
	for word, count := range frequency {
		if count > 1 {
			fmt.Printf("%v -> %d\n", word, count)
		}
	}
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
