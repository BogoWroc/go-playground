package api

type Person struct {
	Weight int
}

type Lift interface {
	Enter(p ...Person)
}

type Display interface {
	Display(message string)
}
