package capstone

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("Point{x=%d, y=%d}", p.x, p.y)
}

type Turtle struct {
	Point
}

func (t Turtle) String() string {
	return fmt.Sprintf("Turtle{%v}", t.Point)
}

func (t *Turtle) Up() {
	t.y++
}

func (t *Turtle) Down() {
	t.x--
}

func (t *Turtle) Left() {
	t.x--
}

func (t *Turtle) Right() {
	t.x++
}
