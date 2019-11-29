package kyu8

func MultipleOfIndex(ints []int) []int {
	var r []int

	for i, v := range ints {
		if i > 0 && v%i == 0 {
			r = append(r, v)
		}
	}
	return r
}
