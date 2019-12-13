package kyu8

func century(year int) int {
	i := int(year / 100)
	if year%100 == 0 {
		return i
	} else {
		return i + 1
	}
}
