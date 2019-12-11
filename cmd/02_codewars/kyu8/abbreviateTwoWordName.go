package kyu8

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	elements := strings.Split(strings.ToUpper(name), " ")
	return fmt.Sprintf("%v.%v", string(elements[0][0]), string(elements[1][0]))
}
