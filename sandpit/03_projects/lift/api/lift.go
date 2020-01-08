package api

type Person struct {
	Weight int
}

type Lift interface {
	IsMaxWeightExceed() bool
	Enter(p ...Person)
}

type Display interface {
	Display(message string)
}
