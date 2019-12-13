package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favoriteIceCreams []string
}

func main() {
	// p1 is not an object it is value of type!
	p1 := person{
		firstName:         "Bob",
		lastName:          "Key",
		favoriteIceCreams: []string{"strawberry"},
	}

	p2 := person{
		firstName:         "Kris",
		lastName:          "House",
		favoriteIceCreams: []string{"blackberry", "chocolate"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	m := map[string]person{
		p1.firstName: p1,
		p2.firstName: p2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
