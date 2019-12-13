package kyu8

import (
	"strings"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	return str == Reverse(str)
}
