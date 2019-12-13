package kyu6

import "strings"

func duplicateCount(s string) int {
	//your code goes here

	s = strings.ToLower(s)
	c := s

	var size = len(s)
	count := 0
	for _, e := range s {
		c = strings.Replace(c, string(e), "", -1)
		if len(c) < size-1 {
			count++
		}
		size = len(c)
	}
	return count //Instead of returning '1', you have to return integer 'i' as answer of solution.
}
