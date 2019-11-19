package capstone

import (
	"encoding/json"
	"fmt"
	"os"
)

// structure attributes must start with capital letter; otherwise json lib will not generate json!!!
type location struct {
	Name string  `json:"name"`
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func PrintLandingSites() {

	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}

	bytes, e := json.MarshalIndent(locations, "", "  ")
	exitOnError(e)
	fmt.Println(string(bytes))
}

// Coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type Coordinate struct {
	D, M, S float64
	H       rune
}

// decimal converts a d/m/s Coordinate to decimal degrees.
func (c Coordinate) decimal() float64 {
	sign := 1.0
	switch c.H {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.D + c.M/60 + c.S/3600)
}

// newLocation from latitude, longitude d/m/s coordinates.
func NewLocation(lat, long Coordinate) location {
	return location{"Location", lat.decimal(), long.decimal()}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
