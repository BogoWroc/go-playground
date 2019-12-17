package kyu7

func WordsToMarks(s string) (sum int) {
	a := 97
	shift := 1
	for _, e := range s {
		sum += int(e) - a + shift
	}
	return
}
