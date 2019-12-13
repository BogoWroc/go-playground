package capstone

import (
	"fmt"
	"math"
)

type Location struct {
	Name string
	Lat  float64
	Long float64
}

func (l location) description() string {
	return fmt.Sprintf("%v %.4f %.4f", l.Name, l.Lat, l.Long)
}

type World struct {
	radius float64
}

func rad(deg float64) float64 { return deg * math.Pi / 180 }

func (w World) distance(p1, p2 Location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	clong := math.Cos(rad(p1.Long - p2.Long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

type GPS struct {
	CurrentLocation     Location
	DestinationLocation Location
	World
}

func (g GPS) distance() float64 {
	return g.World.distance(g.CurrentLocation, g.DestinationLocation)
}

func (g GPS) message() string {
	return fmt.Sprintf("Distance is %.4f km", g.distance())
}

type Rover struct {
}

func (r Rover) Main() {
	world := World{radius: 3389.5}
	gps := GPS{
		CurrentLocation: Location{
			Name: "Bradbury Landing",
			Lat:  -4.5895,
			Long: 137.4417,
		},
		DestinationLocation: Location{
			Name: "Elysium Planitia",
			Lat:  4.5,
			Long: 135.9,
		},
		World: world,
	}

	fmt.Println(gps.message())
}
