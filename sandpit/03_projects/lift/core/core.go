package core

import . "github.com/bogowroc/go-playground/sandpit/03_projects/lift/api"

type lift struct {
	maxWeight int
	persons   []Person
	display   Display
}

func New(maxWeight int, display Display) Lift {
	return &lift{
		maxWeight: maxWeight,
		persons:   make([]Person, 0),
		display:   display,
	}
}

func (l *lift) Enter(p ...Person) {
	l.persons = append(l.persons, p...)
	if l.isMaxWeightExceed() {
		l.display.Display("Max weight exceeded")
	}
}

func (l lift) isMaxWeightExceed() bool {
	if l.calculatePersonsWeight() > l.maxWeight {
		return true
	}

	return false

}

func (l lift) calculatePersonsWeight() (weight int) {
	for _, e := range l.persons {
		weight += e.Weight
	}
	return
}
