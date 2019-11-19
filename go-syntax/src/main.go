package main

import (
	"capstone"
	"chess"
	"fmt"
	"time"
)

func main() {
	lifeExercise()

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