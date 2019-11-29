package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

type square struct {
	l float64
}

func (s square) area() float64 {
	return s.l * s.l
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		l: 5,
	}

	c := circle{
		r: 3,
	}

	info(s)
	info(c)
}
