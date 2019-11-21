package capstone

import "fmt"

type Item struct {
	Name string
}

type Character struct {
	Name     string
	LeftHand *Item
}

func (c *Character) Pickup(i *Item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v picks up a %v\n", c.Name, i.Name)
	c.LeftHand = i
}

func (c *Character) Give(to *Character) {
	if c == nil || to == nil {
		return
	}
	if c.LeftHand == nil {
		fmt.Printf("%v has nothing to give\n", c.Name)
		return
	}
	if to.LeftHand != nil {
		fmt.Printf("%v's hands are full\n", to.Name)
		return
	}
	to.LeftHand = c.LeftHand
	c.LeftHand = nil
	fmt.Printf("%v gives %v a %v\n", c.Name, to.Name, to.LeftHand.Name)
}

func (c Character) String() string {
	if c.LeftHand == nil {
		return fmt.Sprintf("%v is carrying nothing", c.Name)
	}
	return fmt.Sprintf("%v is carrying a %v", c.Name, c.LeftHand.Name)
}
