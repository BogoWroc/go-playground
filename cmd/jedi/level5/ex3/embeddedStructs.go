package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	color  string
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: false,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		color:  "blue",
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)

	fmt.Println(t.color, t.doors, t.fourWheel)
	fmt.Println(s.color, s.vehicle.color, t.doors, t.fourWheel)
}
