package capstone

import "fmt"

type Food string

type Animaler interface {
	move() string
	eat() Food
}

type Dog struct {
	Name string
}

func (d Dog) String() string {
	return fmt.Sprintf("[Dog] My name is %v.", d.Name)
}

func (d Dog) move() string {
	return "Left"
}

func (d Dog) eat() Food {
	return Food("Meat")
}

type Cat struct {
	Name string
}

func (c Cat) String() string {
	return fmt.Sprintf("[Cat] My name is %v.", c.Name)
}

func (c Cat) move() string {
	return "Right"
}

func (c Cat) eat() Food {
	return Food("Fish")
}

func printDetails(a Animaler) {
	fmt.Println(a)
	fmt.Printf("Move -> %v, Eat -> %v\n", a.move(), a.eat())
}

type AnimalsApp struct {
}

func (a AnimalsApp) Main() {
	dog := Dog{Name: "Azor"}
	printDetails(dog)
	cat := Cat{Name: "Filemon"}
	printDetails(cat)
}
