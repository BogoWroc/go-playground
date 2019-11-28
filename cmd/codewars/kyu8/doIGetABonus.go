package main

import (
	"fmt"
)

func BonusTime(salary int, bonus bool) string {
	s := salary

	if bonus {
		s = salary * 10
	}

	return fmt.Sprintf("£%v", s)
}
